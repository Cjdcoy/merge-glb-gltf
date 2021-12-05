package objHandler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
)

// Reader is responsible for writing the Object in a file
type Writer interface {
	Write(object *Object) error
}

func NewWriter(w io.Writer) Writer {
	s := stdWriter{w}

	return s
}

type stdWriter struct {
	w       io.Writer
}

func (e stdWriter) Write(object *Object) error {
	var err error
	if err := e.writeName(object.Name) ; err != nil {
		log.Error(err)
	}
	for _, v := range object.Vertices {
		if err = e.writeVertex(&v); err != nil {
			return err
		}
	}
	for _, t := range object.Textures {
		if err = e.writeTextCoord(&t); err != nil {
			return err
		}
	}
	for _, n := range object.Normals {
		if err = e.writeNormal(&n); err != nil {
			return err
		}
	}
	for _, f := range object.Faces {
		if err = e.writeFace(&f); err != nil {
			return err
		}
	}
	return nil
}

func (e stdWriter) writeName(n string ) error {
	_, err := e.w.Write([]byte("o " + n + "\n"))
	return err
}

func (e stdWriter) writeVertex(v *Vertex) error {
	_, err := e.w.Write([]byte(fmt.Sprintf("v %f %f %f\n", v.X, v.Y, v.Z)))
	return err
}

func (e stdWriter) writeTextCoord(vt *TextureCoord) error {
	_, err := e.w.Write([]byte(fmt.Sprintf("vt %0.3f %0.3f %0.3f\n", vt.U, vt.V, vt.W)))
	return err
}

func (e stdWriter) writeNormal(n *Normal) error {
	_, err := e.w.Write([]byte(fmt.Sprintf("vn %0.4f %0.4f %0.4f\n", n.X, n.Y, n.Z)))
	return err
}

func  (e stdWriter) writeFace(f *Face) error {
	if _, err := e.w.Write([]byte("f ")); err != nil {
		return err
	}
	for idx, p := range f.Points {
		if err := e.writePoint(p); err != nil {
			return err
		}

		if idx != len(f.Points)-1 {
			if _, err := e.w.Write([]byte{' '}); err != nil {
				return err
			}
		}
	}
	if _, err := e.w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

func (e stdWriter) writePoint(p *Point) (err error) {
	if _, err = e.w.Write([]byte(fmt.Sprintf("%d", p.Vertex.Index))); err != nil {
		return
	}

	if p.Texture != nil {
		if _, err = e.w.Write([]byte(fmt.Sprintf("/%d", p.Texture.Index))); err != nil {
			return
		}
	} else if p.Normal != nil {
		if _, err = e.w.Write([]byte("/")); err != nil {
			return
		}
	}

	if p.Normal != nil {
		if _, err = e.w.Write([]byte(fmt.Sprintf("/%d", p.Normal.Index))); err != nil {
			return
		}
	}
	return
}


