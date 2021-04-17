# Release Notes Package *strops* Version 3.0.0

## New Methods
Added methods:
  + StrOps{}.Ptr().JustifyTextInStrField()
  + StrOps{}.Ptr()
  + StrOps{}.Ptr().ReplaceStringChar()
  + StrOps{}.Ptr().RemoveStringChar()

## New Types
Added type:
  + TextJustify - Text Justification Enumeration
  
## More Tests
  + Added tests for new methods.
  + Test coverage is now 93%

## New Architecture
  + The new thread safety protocols will support parallel processing.
    
  + Added 'ePrefix' error prefix string to signatures of all methods returning errors.
    This allows for improved error management and tracking by supplying method chains
    to all returned error messages.

  + Dispersed operational code to supporting library files for better low level access and thread safety.


## Module Requirements
All Version 3+ releases support *Go* modules.
With this release, *go.mod* now supports *Go*
Programming Language Version 1.15.6. 
  