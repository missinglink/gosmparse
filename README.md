# OpenStreetMap PBF Parser in Go

Gosmparse works already, but the API is subject to change.

It has been designed with performance and maximum usage convenience in mind; on an 4 core machine with an SSD it is able to read around 45 MB/s, which would parse a complete planet file in about 12 minutes.