package Compressor

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/* Compress Compress a file with the given parameter of the compression mode
 * @param string : Name of the file to be compressed; string : Name of the compressed new file; int : Compression mode.
 * @return err : error handling
 */
func Compress(Out string, In string, Level int) error {

	output, err := os.Create(Out)
	if err != nil {
		return err
	}
	defer output.Close()

	input, err := os.Open(In)
	if err != nil {
		return err
	}
	defer input.Close()

	buff, err := gzip.NewWriterLevel(output, Level)

	fileinfo, err := input.Stat()
	if err != nil {
		return err
	}

	buff.Name = fileinfo.Name()
	buff.ModTime = fileinfo.ModTime()

	_, err = io.Copy(buff, input)
	if err != nil {
		return err
	}

	if err := buff.Flush(); err != nil {
		return err
	}
	if err := buff.Close(); err != nil {
		return err
	}
	return nil
}

/* Extract extract a file with the given parameter of the compression mode
 * @param string : Name of the file to be compressed; string : Name of the compressed new file; int : Compression mode.
 * @return err : error handling
 */
func Extract(Out string, In string) error {

	output, err := os.Create(Out)
	if err != nil {
		return err
	}
	defer output.Close()

	input, err := os.Open(In)
	if err != nil {
		return err
	}
	defer input.Close()



	buff, err := gzip.NewReader(input)
	if err != nil {
		return err
	}

	fileinfo, err := input.Stat()
	if err != nil {
		return err
	}

	buff.Name = fileinfo.Name()
	buff.ModTime = fileinfo.ModTime()
	_, err = io.Copy(output, buff)
	if err != nil {
		return err
	}

	if err := buff.Close(); err != nil {
		return err
	}
	return nil
}

/* ListDir Reads all files in the given routine and returns a slice containing them
 * @return []string : all files' names
 */
func ListDir(filename string) []string {
	names := make([]string, 0)
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		names = append(names, f.Name())
	}
	return names
}