# SM.MS Util
> A tiny tool to upload images to https://sm.ms

## Usage
1. Single image upload
    ```bash
    $ go run smms.go ~/path/to/image.jpg

    Filename: image.jpg
    Size    : 4669
    W × H   : 144 × 142
    Delete  : https://sm.ms/delete/VYsCNDaQiy35Iw4
    URL     : https://ooo.0o0.ooo/2017/05/27/5928e16ab3424.jpg
    ===========
    ```

2. Folder upload
    ```bash
    $ go run smms.go ~/some/folder/contains/images/

    Filename: foo.jpg
    Size    : 4669
    W × H   : 144 × 142
    Delete  : https://sm.ms/delete/VYsCNDaQiy35Iw4
    URL     : https://ooo.0o0.ooo/2017/05/27/5928e16ab3424.jpg
    ===========
    Filename: bar.jpg
    Size    : 23955
    W × H   : 440 × 394
    Delete  : https://sm.ms/delete/rj4LzfKU1CRsMbD
    URL     : https://ooo.0o0.ooo/2017/05/27/5928e16b41bac.jpg
    ===========
    ```