package controllers

import (
    "log"
    "net/http"

    "music-library/database"
    "music-library/models"
    "music-library/services"

    "github.com/gin-gonic/gin"
)

type SongDetail struct {
    ReleaseDate string `json:"releaseDate"`
    Text        string `json:"text"`
    Link        string `json:"link"`
}

// @Summary Get all songs
// @Description Retrieve a list of songs with optional filtering and pagination
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Filter by group name"
// @Param song query string false "Filter by song name"
// @Param skip query int false "Number of records to skip (default: 0)"
// @Param limit query int false "Number of records to return (default: 10)"
// @Success 200 {array} models.Song "List of songs"
// @Failure 400 {object} models.ErrorResponse "Bad request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /songs [get]
func GetSongs(c *gin.Context) {
    log.Println("[INFO] Starting to fetch songs")
    
    offset := c.DefaultQuery("skip", "0")
    limit := c.DefaultQuery("limit", "10")
    groupFilter := c.Query("group")
    songFilter := c.Query("song")

    log.Printf("[DEBUG] Filters - Group: %s, Song: %s, Offset: %s, Limit: %s", 
        groupFilter, songFilter, offset, limit)

    var songs []models.Song
    query := database.DB

    if groupFilter != "" {
        query = query.Where("\"group\" = ?", groupFilter)
        log.Printf("[DEBUG] Applying group filter: %s", groupFilter)
    }

    if songFilter != "" {
        query = query.Where("song = ?", songFilter)
        log.Printf("[DEBUG] Applying song filter: %s", songFilter)
    }

    if err := query.Offset(offset).Limit(limit).Find(&songs).Error; err != nil {
        log.Printf("[ERROR] Failed to fetch songs: %v", err)
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Database error"})
        return
    }

    log.Printf("[INFO] Successfully fetched %d songs", len(songs))
    c.JSON(http.StatusOK, songs)
}

// @Summary Add a new song
// @Description Add a new song to the library, enriched with details from an external API
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Song details"
// @Success 201 {object} models.Song "Created song"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Failure 503 {object} models.ErrorResponse "External service unavailable"
// @Router /songs [post]
func CreateSong(c *gin.Context) {
    log.Println("[INFO] Starting to create new song")

    var newSong models.Song
    if err := c.ShouldBindJSON(&newSong); err != nil {
        log.Printf("[ERROR] JSON binding failed: %v", err)
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input format"})
        return
    }

    log.Printf("[DEBUG] Received song data - Group: %s, Song: %s", 
        newSong.Group, newSong.Song)

    details, err := services.GetSongDetails(newSong.Group, newSong.Song)
    if err != nil {
        log.Printf("[ERROR] API call failed: %v", err)
        c.JSON(http.StatusServiceUnavailable, models.ErrorResponse{Error: "External service error"})
        return
    }

    log.Printf("[DEBUG] API response: %+v", details)

    // Обработка дефолтных значений
    if details.Text == "" {
        log.Println("[WARN] Empty text field, using default")
        details.Text = "No description available"
    }
    if details.ReleaseDate == "" {
        log.Println("[WARN] Empty release date, using default")
        details.ReleaseDate = "Unknown"
    }
    if details.Link == "" {
        log.Println("[WARN] Empty link, using default")
        details.Link = "No link available"
    }

    newSong.ReleaseDate = details.ReleaseDate
    newSong.Text = details.Text
    newSong.Link = details.Link

    log.Printf("[DEBUG] Final song data before save: %+v", newSong)

    if err := database.DB.Create(&newSong).Error; err != nil {
        log.Printf("[ERROR] Database save failed: %v", err)
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to save song"})
        return
    }

    log.Printf("[INFO] Successfully created song with ID: %d", newSong.ID)
    c.JSON(http.StatusCreated, newSong)
}

// @Summary Delete a song
// @Description Delete a song by its ID
// @Tags Songs
// @Param id path string true "Song ID"
// @Success 200 {object} models.ErrorResponse "Song deleted successfully"
// @Failure 404 {object} models.ErrorResponse "Song not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
    id := c.Param("id")
    log.Printf("[INFO] Starting to delete song with ID: %s", id)

    var song models.Song
    if err := database.DB.Where("id = ?", id).First(&song).Error; err != nil {
        log.Printf("[ERROR] Song not found: %v", err)
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Song not found"})
        return
    }

    log.Printf("[DEBUG] Found song to delete: %+v", song)

    if err := database.DB.Delete(&song).Error; err != nil {
        log.Printf("[ERROR] Delete operation failed: %v", err)
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Deletion failed"})
        return
    }

    log.Printf("[INFO] Successfully deleted song with ID: %s", id)
    c.JSON(http.StatusOK, models.ErrorResponse{Error: "Song deleted successfully"})
}