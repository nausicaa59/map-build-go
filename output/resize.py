import json
import os
import copy
import sys

#----------------------------------
#Define
#----------------------------------
def resize(dim):
	for root, subdirs, files in os.walk("C:/Users/calypso/go/src/github.com/nausicaa59/parse/output/img-map/" + str(dim) + "/"):
		for filename in files:
			pathFile = os.path.join(root, filename).replace("\\","/")
			print(pathFile)
			os.system('magick convert "' + pathFile + '" -resize 256x256 -antialias "'+ pathFile +'"')


"""resize(0)
resize(1)
resize(2)
resize(3)
resize(4)"""

resize(6)