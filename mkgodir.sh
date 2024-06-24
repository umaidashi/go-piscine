#!/bin/sh

PRINTRUNE_FILE=~/etc/ft/printrune.go
SCRIPTNAME=${0}
if [ $# -lt 2 ]
then
	echo "Usage:${SCRIPTNAME} EX_DIRNAME EXCERCISE_NAME [PRINTRUNE_FILE]" 1>&2
	exit
fi
MK_DIRNAME=${1}
if [ -d ${MK_DIRNAME} ]
then
	echo "${MK_DIRNAME} is exist" 1>&2
	exit
fi
EXCERCISE_NAME=${2}.go
if [ -e ${MK_DIRNAME}/vendor/piscine/${EXCERCISE_NAME} ]
then
	echo "${MK_DIRNAME}/vendor/piscine/${EXCERCISE_NAME} is exist" 1>&2
	exit
fi
if [ $# -ge 3 ]
then
	PRINTRUNE_FILE=${3}
fi
if [ ! -e ${PRINTRUNE_FILE} ]
then
	echo "${PRINTRUNE_FILE} is not exist" 1>&2
	exit
fi
mkdir ${MK_DIRNAME}
touch ${MK_DIRNAME}/main.go
mkdir ${MK_DIRNAME}/vendor
mkdir ${MK_DIRNAME}/vendor/ft
cp -p ${PRINTRUNE_FILE} ${MK_DIRNAME}/vendor/ft/
mkdir ${MK_DIRNAME}/vendor/piscine
touch ${MK_DIRNAME}/vendor/piscine/${EXCERCISE_NAME}
