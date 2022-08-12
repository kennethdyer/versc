# versc

Versc is a tool for use in documentation projects that rely on reStructuredText as source code.
It scans the givne source repository, finds all version descriptors that are not in a list 
of supported versions, then prints the finds to stdout in a format similar to `grep`.

## Installation

1. Clone repository:
   ```
   $ cd ~/go/src/github.com/kennethdyer 
   $ git clone git@github.com/kennethdyer/versc
   $ cd versc
   ```
   
2. Install external dependencies:
   ```
   $ go get github.com/spf13/viper
   $ go get github.com/spf13/cobra
   ```
   
4. Install `versc`:

   ```
   $ make
   ```

## Configuration

To configure versc, create a configuration file at `~/.config/avocet/versc.yml`.  The configuration file
should feature a `products` map, which provides specification for each product.    

```yml
products:
  myProductKey:
    name: MyProduct
    versions:
       - "2.1"
       - "2.2"
exclude_pattern:
   - release-notes
```
 
The product specification requires a `name` field, which indicates the likely string to occur 
for the product name.  

It also requires a `versions` field, which is an array of version numbers for the currently supported releases.  Any version that is not found in the array is logged as a match and printed to stdout.

It is best practice to exclude the directory containing release notes, as this is a legitimate place
to mention older versions of a product.

## Usage

To print unsupported version descriptors, call versc with the directory you want to search:

```
$ versc source
```
 
If you have multiple products specified, provide the product key to the search:

```
$ versc source myProductKey
2.0  -  source/reference/feature   Starting in MyProduct 2.0, xyz occurs
```

## Explanation of the Name

The name versc derives from a syncopation of the boring descriptor "version scanner,"
which here has an Old English pronunciation of `/vərʃ/`.

