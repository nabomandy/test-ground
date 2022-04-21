package dao

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"

	geom "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/wkt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type EWKBGeom struct {
	GeomData geom.T
}
type Coord []float64

func NewEWKBGeomPoint(x float64, y float64, srid int) EWKBGeom {
	return EWKBGeom{
		GeomData: geom.NewPoint(geom.XY).MustSetCoords([]float64{x, y}).SetSRID(srid),
	}
}

func NewEWKBGeomLineString(coords []Coord, srid int) EWKBGeom {
	// 배열구조체는 형변환이 안되므로 루프를 돌아서 적용
	coords2 := make([]geom.Coord, 0, len(coords))
	for _, coord := range coords {
		coords2 = append(coords2, geom.Coord(coord))
	}
	return EWKBGeom{
		GeomData: geom.NewLineString(geom.XY).MustSetCoords(coords2).SetSRID(srid),
	}
}

func AppendLineString(line *EWKBGeom, coord Coord) {
	// 배열구조체는 형변환이 안되므로 루프를 돌아서 적용
	got := line.GeomData
	if line == nil {
		got := NewEWKBGeomLineString([]Coord{coord}, 4326)
		line = &got
	}
	switch got.(type) {
	case *geom.LineString:
		// 속성이 LineString 일때만 처리한다
		var l *geom.LineString = got.(*geom.LineString)
		var coords []geom.Coord = l.Coords()
		coords = append(coords, geom.Coord(coord))
		l.SetCoords(coords)
		line.GeomData = l
	default:
		// 넘겨받은 그대로 전달한다
		return
	}
}

// NewEWKBGeomFromText : WKT 문자열 -- POINT(127.323 36.324) -- 을 Geom객체로 전환해 주는 함수
func NewEWKBGeomFromText(WKT string, SRID int) (EWKBGeom, error) {
	got, err := wkt.Unmarshal(WKT)
	if err != nil {
		return EWKBGeom{
			GeomData: nil,
		}, err
	}
	switch got.(type) {
	case *geom.Point:
		got.(*geom.Point).SetSRID(SRID)
	case *geom.LineString:
		got.(*geom.LineString).SetSRID(SRID)
	case *geom.Polygon:
		got.(*geom.Polygon).SetSRID(SRID)
	case *geom.MultiPoint:
		got.(*geom.MultiPoint).SetSRID(SRID)
	case *geom.MultiLineString:
		got.(*geom.MultiLineString).SetSRID(SRID)
	case *geom.MultiPolygon:
		got.(*geom.MultiPolygon).SetSRID(SRID)
	case *geom.GeometryCollection:
		got.(*geom.GeometryCollection).SetSRID(SRID)
	}
	return EWKBGeom{
		GeomData: got,
	}, nil
}

func (g EWKBGeom) MarshalJSON() ([]byte, error) {
	got, err := wkt.Marshal(g.GeomData)
	if err != nil {
		got = ""
	}
	return json.Marshal(got)
}

func (g *EWKBGeom) UnmarshalJSON(b []byte) error {
	gt, err := ewkb.Unmarshal(b)
	if err != nil {
		return err
	}
	*g = EWKBGeom{
		GeomData: gt,
	}
	return nil
}

func (g EWKBGeom) String() string {
	srid := g.GeomData.SRID()
	got, err := wkt.Marshal(g.GeomData)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("SRID=%d;%s", srid, got)
}

func (g EWKBGeom) SRID() int {
	return g.GeomData.SRID()
}

func (g *EWKBGeom) Scan(input interface{}) error {
	if input != nil {
		dst, err := hex.DecodeString(input.(string))
		if err != nil {
			return err
		}
		gt, err := ewkb.Unmarshal(dst)
		if err != nil {
			return err
		}
		*g = EWKBGeom{
			GeomData: gt,
		}
	}
	return nil
}

func (g EWKBGeom) Value() (driver.Value, error) {
	srid := g.GeomData.SRID()
	got, err := wkt.Marshal(g.GeomData)
	if err != nil {
		return nil, err
	}
	value := fmt.Sprintf("SRID=%d;%s", srid, got)
	return value, nil
}

// GormDataType gorm common data type
func (EWKBGeom) GormDataType() string {
	return "geometry"
}

// GormDBDataType gorm db data type
func (EWKBGeom) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "postgres":
		return "geometry"
	}
	return ""
}
