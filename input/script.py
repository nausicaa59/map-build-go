import os
import shutil
import sys
from selenium import webdriver
from selenium.webdriver.common.keys import Keys

PATH_DRIVER_GOOGLE = "C:/piloteweb/chromedriver.exe"
basePath = "input/"
pathDirInput 	= basePath + "listePseudo/"
pathDirD3 		= basePath + "d3/"
pathDirOutput 	= basePath + "svg/"
driver 			= webdriver.Chrome(PATH_DRIVER_GOOGLE)

for f in os.listdir(pathDirInput):
	name 			= f.split(".")[0]
	pathFileInput 	= pathDirInput + f
	pathFileD3 		= pathDirD3 + "input.csv"
	pathFileOutput 	= pathDirOutput + name + ".xml"
	shutil.copy(pathFileInput, pathFileD3)

	driver.get("http://d3generate.dev/")
	html = driver.find_element_by_css_selector("svg").get_attribute('outerHTML')
	with open(pathFileOutput, 'w') as f:
		f.write(html)


driver.close()
