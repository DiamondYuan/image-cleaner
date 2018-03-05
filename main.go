package main

import (
	"github.com/docker/docker/client"
	"context"
	"github.com/docker/docker/api/types"
	"time"
	"strings"
	"bufio"
	"os"
	"regexp"
	"flag"
	"log"
)

var cli, _ = client.NewEnvClient()
var ctx = context.Background()

const (
	defaultWaitTime = 5;
)

type Image struct {
	Name     string
	RepoTags string
	Created  time.Time
	Tag      string
	ID       string
}

var dryRun bool
var notWait bool
var infinite  bool
var whiteList []string

func init() {
	flag.BoolVar(&dryRun, "dryRun", false, "just list containers to remove")
	flag.BoolVar(&notWait, "notWait", false, "Don't Wait")
	flag.BoolVar(&infinite, "infinite", false, "Infinite run")
	flag.Parse()
	if os.Getenv("DOCKER_HOST") == "" {
		err := os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
		if err != nil {
			log.Fatalf("error setting default DOCKER_HOST: %s", err)
		}
	}
	whiteList, _ = ReadConfig("whiteList")
}

func main() {
	clean()
	if !infinite {
		os.Exit(0)
	}
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			clean()
		}
	}()
	select {}

}

func clean()  {
	log.Printf("以下镜像会被删除")
	list, _ := cli.ImageList(ctx, types.ImageListOptions{
		All: true,
	})
	var imagesToClean []Image
	for _, value := range list {
		image := Image{
			ID:       value.ID,
			Tag:      strings.Split(value.RepoTags[0], ":")[1],
			Name:     strings.Split(value.RepoTags[0], ":")[0],
			RepoTags: value.RepoTags[0],
		}
		if !inWhiteList(image) {
			imagesToClean = append(imagesToClean, image)
		}
	}
	for _, node := range imagesToClean {
		log.Printf("%s   %s", node.RepoTags, node.ID)
	}

	if dryRun {
		log.Printf("程序完毕 正常退出")
		return
	}
	if !notWait {
		log.Printf("等待 %d 秒 此时你可以取消", defaultWaitTime)
		time.Sleep(time.Duration(defaultWaitTime) * time.Second)
	}
	log.Printf("开始清理")
	for _, node := range imagesToClean {
		log.Printf("Clean %s   %s", node.RepoTags, node.ID)
		_, err := cli.ImageRemove(ctx, node.ID, types.ImageRemoveOptions{
			Force:         false,
			PruneChildren: false,
		})
		if err != nil {
			log.Println(err)
		}
	}
}








func inWhiteList(image Image) bool {
	for _, node := range whiteList {
		r, _ := regexp.Compile(node)
		findString := r.FindString(image.RepoTags);
		if findString == image.RepoTags {
			return true
		}
	}
	return false
}

func ReadConfig(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
