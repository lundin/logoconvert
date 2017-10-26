// mytest.go
package main

import (
	_"fmt"
	"io"	
	"path/filepath"
	"io/ioutil"
	"net/http"
	"image/png"
	"github.com/jsummers/gobmp"
	    "image" 
	"image/gif"
	"image/jpeg"
	  "image/color" 
    "image/draw" 
	"gopkg.in/cheggaaa/pb.v1"
	_"golang.org/x/image/bmp"
	"os"
	"log"
	"strconv"
	"encoding/json"
	"fmt"
)


type ImageInfo struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Type        string `json:"type"`
	Space       string `json:"space"`
	Alpha       bool   `json:"hasAlpha"`
	Profile     bool   `json:"hasProfile"`
	Channels    int    `json:"channels"`
	Orientation int    `json:"orientation"`
}

func readFile(file string) io.Reader {
	buf, _ := os.Open(file)
	return buf
}



 func createDirectory(dirName string) bool {
         src, err := os.Stat(dirName)

         if os.IsNotExist(err) {
                 errDir := os.MkdirAll(dirName, 0755)
                 if errDir != nil {
                         panic(err)
                 }
                 return true
         }

         if src.Mode().IsRegular() {
                 fmt.Println(dirName, "already exist as a file!")
                 return false
         }

         return false
 }

func main() {

fmt.Print("Converting nasty BMP to JPG...wait..and chill \n")	 

pwd, err := os.Getwd()
    		if err != nil {
        	fmt.Println(err)
        	os.Exit(1)
    		}


createDirectory(pwd+"/done")
createDirectory(pwd+"/done/logo")

count := 0
files, _ := ioutil.ReadDir(pwd)
 var opt jpeg.Options

         opt.Quality = 100
			for _, f := range files {
if filepath.Ext(f.Name()) == ".PNG" || filepath.Ext(f.Name()) == ".JPG" || filepath.Ext(f.Name()) == ".JPEG" || filepath.Ext(f.Name()) == ".jpg" || filepath.Ext(f.Name()) == ".jpeg" || filepath.Ext(f.Name()) == ".png" || filepath.Ext(f.Name()) == ".bmp" || filepath.Ext(f.Name()) == ".BMP" || filepath.Ext(f.Name()) == ".GIF" || filepath.Ext(f.Name()) == ".gif" {
		count++	
}
if filepath.Ext(f.Name()) == ".BMP" || filepath.Ext(f.Name()) == ".bmp" {

	 file, err := os.Open(f.Name())
 if err != nil {
  fmt.Print(err)
 }

 defer file.Close()

img, err := gobmp.Decode(file)
 if err != nil {
  fmt.Print(err)
  
 } else {


 out, err := os.Create(f.Name()[0:len(f.Name())-len(filepath.Ext(f.Name()))]+".jpg")
 if err != nil {
  fmt.Println(err)
  
 }
 err = jpeg.Encode(out, img,&opt)
 if err != nil {
  fmt.Println(err)
  
 }

 }


}




if filepath.Ext(f.Name()) == ".GIF" || filepath.Ext(f.Name()) == ".gif" {

	 file, err := os.Open(f.Name())
 if err != nil {
  fmt.Print(err)
 }

 defer file.Close()

img, err := gif.Decode(file)
 if err != nil {
  fmt.Print(err)
  
 } else {


 out, err := os.Create(f.Name()[0:len(f.Name())-len(filepath.Ext(f.Name()))]+".jpg")
 if err != nil {
  fmt.Println(err)
  
 }
 err = jpeg.Encode(out, img,&opt)
 if err != nil {
  fmt.Println(err)
  
 }

 }


}





if filepath.Ext(f.Name()) == ".PNG" || filepath.Ext(f.Name()) == ".png" {

	 file, err := os.Open(f.Name())
 if err != nil {
  fmt.Print(err)
 }

 defer file.Close()

src, err := png.Decode(file)
 if err != nil {
  fmt.Print(err)
  
 } else {


    backgroundColor := color.RGBA{0xFF, 0xFF, 0xFF, 0xff} // Dark red. 

 dst := image.NewRGBA(src.Bounds()) 
    draw.Draw(dst, dst.Bounds(), image.NewUniform(backgroundColor), 
image.Point{}, draw.Src) 
    draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Over) 

    dstFile, err :=os.Create(f.Name()[0:len(f.Name())-len(filepath.Ext(f.Name()))]+".jpg")
    if err != nil { 
        log.Fatal(err) 
    } 
    defer dstFile.Close() 
    err = jpeg.Encode(dstFile, dst, &opt) 
    if err != nil { 
        log.Fatal(err) 
    } 



 }





}

			}

fmt.Print("Done...starting resize... \n")
bar := pb.StartNew(count)

files, _ = ioutil.ReadDir(pwd)
			for _, f := range files {
				
if filepath.Ext(f.Name()) == ".JPG" || filepath.Ext(f.Name()) == ".JPEG" || filepath.Ext(f.Name()) == ".jpg" || filepath.Ext(f.Name()) == ".jpeg" {
//fmt.Print(f.Name()+"\n")

//method:="resize"

		buf := readFile(f.Name())

			url := "http://imageutil.partsworkshop.com/info"
			//url := "http://localhost:8080/extract?top=400&left=100&areawidth=800&areaheight=720"

			res, err := http.Post(url, "image/jpeg", buf)
			if err != nil {
				log.Fatal("Cannot perform the request")
			}

			if res.StatusCode != 200 {
				log.Fatalf("Invalid response status: %s", res.Status)
			}

			image, err := ioutil.ReadAll(res.Body)
			Myimg := ImageInfo{}
			if err := json.Unmarshal(image, &Myimg); err != nil {
				panic(err)
			}

		
			

	
	
	
	




		//create large thumb

if (Myimg.Width<100 && Myimg.Height<39) {
			//method:="enlarge"
			
			if (Myimg.Width<Myimg.Height) {
				//fmt.Print(Myimg.Height)
				//fmt.Print(float64(Myimg.Height)/float64(Myimg.Width))
			height:=(float64(Myimg.Height)/float64(Myimg.Width))*100
			width:=100.000
			if (height<39) {
							height=39
							width=(float64(Myimg.Height)/float64(Myimg.Width))*39
						}
			//fmt.Printf("i am height %f ",(height))
			url = "http://imageutil.partsworkshop.com/enlarge?width="+strconv.Itoa(int(width))+"&height="+strconv.Itoa(int(height))+"&quality=100"	
		} else {
//fmt.Print("now in here")
			width:=(float64(Myimg.Width)/float64(Myimg.Height))*100
			height:=39.000
						//fmt.Printf("i am width %f ",(width))
						if (width<100) {
						width=100
						height=100/(float64(Myimg.Width)/float64(Myimg.Height))
						}
			url = "http://imageutil.partsworkshop.com/enlarge?width="+strconv.Itoa(int(width))+"&height="+strconv.Itoa(int(height))+"&quality=100"
			}
			

	buf = readFile(f.Name())
	//url = "http://imageutil.partsworkshop.com/enlarge?width=1000&height=1810&quality=100"
	//defer ts.Close()

	res, err = http.Post(url, "image/jpeg", buf)
	if err != nil {
		//t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		//t.Fatalf("Invalid response status: %s", res.Status)
	}

	image, err = ioutil.ReadAll(res.Body)
	if err != nil {
		//t.Fatal(err)
	}
	if len(image) == 0 {
		//t.Fatalf("Empty response body")
	}

	err = ioutil.WriteFile(pwd+"/"+f.Name(), image, 0644)
				if err != nil {
					panic(err)
				}
}



	buf = readFile(pwd+"/"+f.Name())
	url = "http://imageutil.partsworkshop.com/resize?width=100&height=39&quality=100&nocrop=true&extend=background&background=255,255,255&embed=true&background=255,255,255"
	//defer ts.Close()

	res, err = http.Post(url, "image/jpeg", buf)
	if err != nil {
		//t.Fatal("Cannot perform the request")
	}

	if res.StatusCode != 200 {
		//t.Fatalf("Invalid response status: %s", res.Status)
	}

	image, err = ioutil.ReadAll(res.Body)
	if err != nil {
		//t.Fatal(err)
	}
	if len(image) == 0 {
		//t.Fatalf("Empty response body")
	}

	err = ioutil.WriteFile(pwd+"/done/logo/"+f.Name()[0:len(f.Name())-len(filepath.Ext(f.Name()))]+".jpg", image, 0644)
				if err != nil {
					panic(err)
				}




bar.Increment()

			}

		}
		bar.FinishPrint("The End!")
/*
 buf = readFile("testpontus_medium.jpg")
        url = "http://imageutil.partsworkshop.com/thumbnail?extend=background&background=255,255,255"
        //defer ts.Close()

        res, err = http.Post(url, "image/jpeg", buf)
        if err != nil {
                //t.Fatal("Cannot perform the request")
        }

        if res.StatusCode != 200 {
                //t.Fatalf("Invalid response status: %s", res.Status)
        }

        image, err = ioutil.ReadAll(res.Body)
        if err != nil {
                //t.Fatal(err)
        }
        if len(image) == 0 {
                //t.Fatalf("Empty response body")
        }


   err = ioutil.WriteFile("testpontus_medium2.jpeg", image, 0644)
                                if err != nil {
                                        panic(err)
                                }
*/
}
