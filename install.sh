copy_default_config_folder() {
    if [ ! -d "$HOME/.theeman" ]; then
        cp -r .theeman/ $HOME/.theeman
    fi
}

build_go_project() {
    go install .
}

copy_default_config_folder
build_go_project
