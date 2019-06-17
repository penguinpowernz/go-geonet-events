package events

import (
	"testing"

	"github.com/penguinpowernz/go-geonet"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProcessor(t *testing.T) {
	Convey("given a processor", t, func() {
		p := NewProcessor()

		Convey("when a new quake is passed", func() {
			qk := geonet.Quake{}
			qk.PublicID = "1"
			qks := []geonet.Quake{qk}
			evts := p.Process(qks)

			Convey("then an event should be returned", func() {
				So(len(evts), ShouldEqual, 1)
			})

			Convey("then it should have the type new", func() {
				So(evts[0].Type, ShouldEqual, "new")
			})

			Convey("then the cache should contain one item", func() {
				So(p.cache.ItemCount(), ShouldEqual, 1)
			})

			Convey("and it is updated", func() {
				qk.Depth = 5.5
				qks := []geonet.Quake{qk}
				evts := p.Process(qks)

				Convey("then an event should be returned", func() {
					So(len(evts), ShouldEqual, 1)
				})

				Convey("then it should have the type updated", func() {
					So(evts[0].Type, ShouldEqual, "updated")
				})

				Convey("then it should contain the fields that were updated", func() {
					So(evts[0].UpdatedFields, ShouldContain, "depth")
				})

				Convey("then the cache should contain one item", func() {
					So(p.cache.ItemCount(), ShouldEqual, 1)
				})

			})
		})
	})
}
