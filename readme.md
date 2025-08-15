# litxap-fwew

A small wrapper to add support for the [Fwew Na'vi dictionary library](https://github.com/fwew/fwew-lib) to Litxap. 
It was previously part of `litxap-service`, but I moved it out to make integrating Litxap into Sarfya easier.

The `Global` function will return the shared fwew instance, and there is a `MultiWordPartDictionary` helper to put all those half-words like "tsaheyl" into a dummy dictionary since litxap handles each word separately.

