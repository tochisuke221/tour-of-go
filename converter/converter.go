package converter

import (
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "strings"

    "golang.org/x/image/bmp"
    "golang.org/x/image/tiff"
    "golang.org/x/image/webp"
)

// Convert converts an image from srcPath to dstPath with specified formats.
func Convert(srcPath, dstPath, srcFormat, dstFormat string) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return fmt.Errorf("failed to open source file: %v", err)
    }
    defer srcFile.Close()

    var img image.Image
    switch strings.ToLower(srcFormat) {
    case "jpg", "jpeg":
        img, err = jpeg.Decode(srcFile)
    case "png":
        img, err = png.Decode(srcFile)
    case "bmp":
        img, err = bmp.Decode(srcFile)
    case "tiff":
        img, err = tiff.Decode(srcFile)
    case "webp":
        img, err = webp.Decode(srcFile)
    default:
        return fmt.Errorf("unsupported source format: %s", srcFormat)
    }
    if err != nil {
        return fmt.Errorf("failed to decode source file: %v", err)
    }

    dstFile, err := os.Create(dstPath)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %v", err)
    }
    defer dstFile.Close()

    switch strings.ToLower(dstFormat) {
    case "jpg", "jpeg":
        err = jpeg.Encode(dstFile, img, nil)
    case "png":
        err = png.Encode(dstFile, img)
    case "bmp":
        err = bmp.Encode(dstFile, img)
    case "tiff":
        err = tiff.Encode(dstFile, img, nil)
    default:
        return fmt.Errorf("unsupported destination format: %s", dstFormat)
    }
    if err != nil {
        return fmt.Errorf("failed to encode destination file: %v", err)
    }

    return nil
}
