#!/bin/bash
if [[ $EUID -eq 0 ]]; then
  echo "You shouldn't run this script as root"
  exit 1
fi

readonly APP_NAME='Notes Manager'
readonly SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd ) # Get script dir

# Load config
source $SCRIPT_DIR/config

handleUnits() {
    units_tmp=$1/*.$NOTEBOOK_EXT
    units=()
    for f in $units_tmp; do
        units+=(`basename ${f%.*}`)
    done

    finished=0
    while [ $finished -ne 1 ]; do
        out=$(zenity --list --title="Choose a unit - ${APP_NAME}" --column="Unit" ${units[@]})
        if [ -z $out ]; then
            return
        fi

        tema="$1/$out.$NOTEBOOK_EXT"
        $NOTEBOOK_BIN $tema
    done
}

handleSubjects() {
    finished=0
    while [ $finished -ne 1 ]; do
        sub_tmp=$SCRIPT_DIR/data/*/
        sub=()
        for f in $sub_tmp; do
            sub+=(`basename $f`)
        done

        out=$(zenity --list --title="Choose a subject ${APP_NAME}" --column="Subject" ${sub[@]} '/--/ New /--/')
        if [[ -z $out ]]; then
            return
        fi

        if [[ $out == '/--/ New /--/' ]]; then
            newFolder=`zenity --entry --title="New subject" --text="Write the name of your new subject:"`
            if [[ -n $newFolder ]]; then
                mkdir $SCRIPT_DIR/data/$newFolder
            fi
        else
            units_dir="$SCRIPT_DIR/data/$out"
	        if [ "$(ls -A $units_dir)" ]; then
                handleUnits $units_dir
	        else
                zenity --error --text="That subject does not have units inside"
            fi
        fi
    done
}

main() {
    finished=0
    while [ $finished -ne 1 ]; do
        out=$(zenity --list --title="Main - ${APP_NAME}" --column="ID" --column="Section" 0 Subjects)
        if [ -z $out ]; then
            finished=1
            echo "Bye bye"
            exit 0
        fi

        case $out in
            0) handleSubjects;;
        esac
    done
}

main
