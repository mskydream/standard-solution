package main

import (
	"archive/tar"
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	writeTar()
	readTar()
	writeZip()
	readZip()
	noCompression()
}

var files = []string{
	"proverbs1.txt",
	"proverbs2.txt",
	"proverbs2.txt",
}

func writeTar() {
	tf, err := os.Create("proverbs.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer tf.Close()

	tw := tar.NewWriter(tf)

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		s, _ := f.Stat()
		h := &tar.Header{
			Name:    s.Name(),
			Mode:    0666,
			Size:    s.Size(),
			ModTime: s.ModTime(),
		}

		if err := tw.WriteHeader(h); err != nil {
			log.Fatalln(err)
		}

		b, err := io.ReadAll(f)
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := tw.Write(b); err != nil {
			log.Fatalln(err)
		}
	}
}

func readTar() {
	f, err := os.Open("proverbs.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	tr := tar.NewReader(f)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			fmt.Println("--- --- Reach the end")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s--- --- --- --- ---\n", h.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
	}
}

func writeZip() {
	zf, err := os.Create("proverbs.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer zf.Close()

	zw := zip.NewWriter(zf)

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		fw, err := zw.Create(f.Name())
		if err != nil {
			log.Fatalln(err)
		}

		b, err := io.ReadAll(f)
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := fw.Write(b); err != nil {
			log.Fatalln(err)
		}
	}

	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}

func readZip() {
	r, err := zip.OpenReader("proverbs.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("%s--- --- --- --- ---\n", f.Name)

		rc, err := f.Open()
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := io.Copy(os.Stdout, rc); err != nil {
			log.Fatalln(err)
		}

		rc.Close()
	}
}

func noCompression() {
	zf, err := os.Create("proverbs-nocompress.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer zf.Close()

	zw := zip.NewWriter(zf)

	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.NoCompression)
	})

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		_, err = zw.Create(f.Name())
		if err != nil {
			log.Fatalln(err)
		}
	}
}
