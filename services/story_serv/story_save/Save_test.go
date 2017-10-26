package story_save

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
	"reflect"
)

func TestGormBookRepository(t *testing.T) {
	Convey("Given a valid JSON string in a Save struct", t, func(){
		validJsonString := `{
								"Name" : "Martha Stewart",
								"Class" : "Archer",
								"HP"    : 50,
								"Inventory" : {
									"minor_potion_healing" : { "quantity" : 1 }
								},
								"Morale" : 100
							}`

		initialSaveData := Save{JsonString:validJsonString}

		Convey("Creating an encoded save string should work", func (){
			encodedSaveString, err := initialSaveData.EncodedSaveString()

			So(err,ShouldBeNil)
			So(len(encodedSaveString),ShouldBeGreaterThan, 0)

			Convey("Given a valid encoded save string, decoding it to a Save struct should work", func (){
				decodedSave, err := DecodeSaveString(encodedSaveString)

				So(err, ShouldBeNil)

				Convey("The decoded save struct should have a valid JsonString that has all the same values as before", func (){
					var person1 interface{}
					var person2 interface{}

					unmarshallErr1 := json.Unmarshal([]byte(initialSaveData.JsonString), &person1)
					unmarshallErr2 := json.Unmarshal([]byte(decodedSave.JsonString), &person2)

					So(unmarshallErr1, ShouldBeNil)
					So(unmarshallErr2, ShouldBeNil)

					So(reflect.DeepEqual(person1,person2), ShouldBeTrue)
				})
			})

		})



	})

}