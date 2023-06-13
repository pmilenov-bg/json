# Traverse local file system

The goal of this program is to learn testing and working with jsons.

## Goal 1
Traverse the file system and display all files and folders in a certain path.

### DOD (definition of done)

You have passing tests.

Hints:
Learn how to work and create temp files and folders.

Create a test which will have its "Before" Method
In "Before" - Withing that temp folder location create folders pete/surf
Create empty files in pete/cv.txt and pete/surf/general.txt

Finally run the test asserting you have found the general.txt cv.txt and the two folders.

On tear down, remove all files and folders.

## Goal 2

Export the list of the files in flat: json

    [
        {filepath: 'cv, kind: 'file'}
        {filepath: 'pete/surf/general.txt, kind: 'file'}
    ]
    

### DOD

You have a test validating the output json


## Goal 3
Working with Interfaces


    [
        {filepath: 'pete, kind: 'folder'},
        {filepath: 'pete/surf, kind: 'folder'},
        {filepath: 'cv, kind: 'file'}
        {filepath: 'pete/surf/general.txt, kind: 'file'}
    ]
