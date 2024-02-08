package database
 
type Article struct {
   ID          string `json:"id"`
   Title       string `json:"title"`
   Author      string `json:"author"`
   Description string `json:"description"`
   Rate        int    `json:"rate"`
}
