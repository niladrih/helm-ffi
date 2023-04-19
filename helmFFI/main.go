package main

import "C"
import (
	"fmt"
	"io"
	// "os"
	"path/filepath"

	// "helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/helmpath"
)


// type createOptions struct {
// 	starter    string
// 	name       string
// 	starterDir string
// }



/** func (o *createOptions) run(out io.Writer) error {
	fmt.Fprintf(out, "Creating %s\n", o.name)
	fmt.Println(o.name)
	fmt.Println(o.starter)
	fmt.Println(o.starterDir)
	chartname := filepath.Base(o.name)
	cfile := &chart.Metadata{
		Name:        chartname,
		Description: "A Helm chart for Kubernetes",
		Type:        "application",
		Version:     "0.1.0",
		AppVersion:  "0.1.0",
		APIVersion:  chart.APIVersionV2,
	}

	if o.starter != "" {
		// Create from the starter
		lstarter := filepath.Join(o.starterDir, o.starter)
		// If path is absolute, we don't want to prefix it with helm starters folder
		if filepath.IsAbs(o.starter) {
			lstarter = o.starter
		}
		return chartutil.CreateFrom(cfile, filepath.Dir(o.name), lstarter)
	}

	chartutil.Stderr = out
	_, err := chartutil.Create(chartname, filepath.Dir(o.name))
	return err
}**/

func CreateHelm(out io.Writer, args []string) *C.char {
	// o := &createOptions{}
	name := args[0]
	starterDir := helmpath.DataPath("starters")
	fmt.Fprintf(out, "Creating %s\n", name)
	starter :=  ""
	chartname := filepath.Base(name)
	cfile := &chart.Metadata{
		Name:        chartname,
		Description: "A Helm chart for Kubernetes",
		Type:        "application",
		Version:     "0.1.0",
		AppVersion:  "0.1.0",
		APIVersion:  chart.APIVersionV2,
	}

	if starter != "" {
		// Create from the starter
		lstarter := filepath.Join(starterDir, starter)
		// If path is absolute, we don't want to prefix it with helm starters folder
		if filepath.IsAbs(starter) {
			lstarter = starter
		}
		result := "error: " + chartutil.CreateFrom(cfile, filepath.Dir(name), lstarter).Error()
		return C.CString(result)
		// return chartutil.CreateFrom(cfile, filepath.Dir(name), lstarter)
	}

	chartutil.Stderr = out
	_, err := chartutil.Create(chartname, filepath.Dir(name))
	// return err
	// return o.run(out)
	result := "error: " + err.Error()
    return C.CString(result)
}

func main() {}
// func main() {
// 	out := os.Stdout
// 	if err := CreateHelm(out, []string{"sample"}); err != nil {
// 		fmt.Println("error in creating the chart: ", err)
// 		return
// 	}
// }