package entities

type IDogs struct {
    ID        string `json:"id,omitempty" bson:"_id,omitempty"`
    Name      string `json:"name" bson:"name"`
    Breed     string `json:"breed" bson:"breed"`
    Age       int    `json:"age" bson:"age"`
    IsGoodBoy bool   `json:"isGoodBoy" bson:"isGoodBoy"`
}