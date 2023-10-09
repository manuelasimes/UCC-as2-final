package cache 

import (
	"github.com/bradfitz/gomemcache/memcache"
	"fmt"
    e "user-res-api/utils/errors"
    json "github.com/json-iterator/go"
    "user-res-api/dto"
	log "github.com/sirupsen/logrus"
 
)

var (
    cacheClient *memcache.Client
)

func Init_cache() {
	cacheClient = memcache.New("localhost:11211")
    fmt.Println("Initialized cache", cacheClient)
    log.Info("Initialized cache")
}

func Set(key string, value []byte) {

    //key := createCacheKey(id, startDate)
    //key := strconv.Itoa(id) + strconv.Itoa(startDate)
    if err := cacheClient.Set(&memcache.Item{
        Key: key, 
        Value: value,
    }); err != nil {
        fmt.Println("Error setting item in cache", err)
    }

}


func Get(key string)  (dto.Availability, e.ApiError){
    fmt.Println("entro")
  
    fmt.Println("paso la linea 34")
    response, err := cacheClient.Get(key)
    fmt.Println("paso la linea 36")
    if err != nil {
        if err == memcache.ErrCacheMiss {
            return dto.Availability{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", key))
        }
        errorMsg := fmt.Sprintf("Error getting item from cache: %s", key)
        fmt.Println(errorMsg)
        return dto.Availability{},  e.NewInternalServerApiError(errorMsg, err)
    }
    var responseDto dto.Availability 
    if err := json.Unmarshal(response.Value, &responseDto); err != nil {
        return dto.Availability{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", key), err)

    }
    return responseDto, nil
}

// func createCacheKey(id int, startDate int) string {
//     return fmt.Sprintf("reservation:%d:%d", id, startDate)
// }
// func main() {
//     Init_cache()
//     value := []byte("some data")
//     Set(1, 20231009, value)
//     result := Get(1, 20231009)
//     fmt.Printf("Result: %s\n", string(result))
// }