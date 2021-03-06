package extractor

import "time"

type Overview struct {
	XSize int32 `json:"x_size"`
	YSize int32 `json:"y_size"`
}

type GeoMetaData struct {
	DataSetName  string      `json:"ds_name"`
	NameSpace    string      `json:"namespace,omitempty"`
	Type         string      `json:"array_type"`
	RasterCount  int32       `json:"raster_count"`
	TimeStamps   []time.Time `json:"timestamps"`
	Heights      []float64   `json:"heights,omitempty"`
	Overviews    []*Overview `json:"overviews,omitempty"`
	XSize        int32       `json:"x_size"`
	YSize        int32       `json:"y_size"`
	GeoTransform []float64   `json:"geotransform"`
	Polygon      string      `json:"polygon"`
	ProjWKT      string      `json:"proj_wkt"`
	Proj4        string      `json:"proj4"`
	Min          float64     `json:"min"`
	Max          float64     `json:"max"`
	Mean         float64     `json:"mean"`
	StdDev       float64     `json:"stddev"`
}

type GeoFile struct {
	FileName string         `json:"filename,omitempty"`
	Driver   string         `json:"file_type"`
	DataSets []*GeoMetaData `json:"geo_metadata"`
}

type POSIXDescriptor struct {
	GID   uint32 `json:"gid"`
	Group string `json:"group"`
	UID   uint32 `json:"uid"`
	User  string `json:"user"`
	Size  int64  `json:"size"`
	Mode  string `json:"mode"`
	Type  string `json:"type"`
	INode uint64 `json:"inode"`
	MTime int64  `json:"mtime"`
	ATime int64  `json:"atime"`
	CTime int64  `json:"ctime"`
}
