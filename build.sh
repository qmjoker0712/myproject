#!/bin/bash
# --------------------------------------------------
current_path=$(cd `dirname $0`; pwd)
###AUTOTEST目录指定，使用前需要添加并配置文件：../uni-autotest/build_tool/conf/${PROCESSNAME}_conf.sh
PROCESS_NAME=myproject
AUTOTEST_TOOL_PATH=${current_path}/../uni-autotest/build_tool
AUTOTEST_GIT_PATH=http://code.zsbatech.com/uniautotest/uni-autotest.git
function echo_red
{
    local content=$@
    echo -e "\033[32m $content \033[0m"
    return 0
}

function echo_green
{
    local content=$@
    echo -e "\033[32m $content \033[0m"
    return 0
}

function exec_func(){
    eval $@
    [[ $? -ne 0 ]] && {
        echo_red "cmd[$@] execute fail!"
        exit 1
    }
    return 0
}

##执行uni_autotest/build_tool/main.sh中的命令
## $1 => command
function exec_command(){
    if [[ -z $1 ]];then
        echo_red "$FUNCNAME param is null!"
        exit 1
    fi
    cd ${AUTOTEST_TOOL_PATH} >/dev/null 2>&1
    chmod +x ./main.sh 2>/dev/null
    ./main.sh $@
    [[ $? -ne 0 ]] && {
        echo_red "$FUNCNAME $@ execute fail!"
        cd - >/dev/null 2>&1
        exit 1
    }
    cd - >/dev/null 2>&1
    return 0
}

##初始化uni_autotest工具目录
function init_autotest_env(){
    if [[ -d ${current_path}/../uni-autotest ]];then
        rm -rf ${current_path}/../uni-autotest 2>/dev/null
    fi
    cd ${current_path}/../
    rm -rf uni-autotest.tar.gz 2>/dev/null
    mkdir uni-autotest
    wget  http://117.50.18.210:8089/uni-autotest_release/uni-autotest.tar.gz >/dev/null 2>&1
    local ret_wget=$?
    tar -zxvf uni-autotest.tar.gz  -C ./uni-autotest >/dev/null 2>&1
    local ret_tar=$?
    if [[ $ret_wget -ne 0 || $ret_tar -ne 0 ]];then
        cd -
        echo "autotest env init fail!"
        return 1
    fi
    cd -
    cp -rf ${current_path}/../uni-autotest/build_tool/conf/${PROCESS_NAME}_conf.sh ${current_path}/../uni-autotest/build_tool/conf.sh >/dev/null 2>&1
    return 0
}


## $1 => build param, debug | release
function build() {
    make all
    if [[ $? -ne 0 ]];then
        echo_red "make all fail"
        return 1
    fi
    local output_path=${current_path}/output
    mkdir -p ${output_path}/bin
    cp ${current_path}/build/bin/* ${output_path}/bin/    

    return 0
}

function usage
{
    echo_green "
Usage:
    $0 [\$1]
Options:
    h|-h|help|-help   usage help
    init_env          初始换go语言环境，并配置环境变量
    buildd            编译整个项目(debug)。
    buildr            编译整个项目(release)。    
    package           产出打包
                      目录结构为：
                      xxxx.tar.gz ─┬conf
                                   ├bin
                                   ├data
                                   ├log
    vet_local         静态代码语法错误检查，针对local测试
    vet_master        静态代码语法错误检查，针对master测试
    ut_coverage       全量UTcase 覆盖率测试，并生成覆盖率报告
    local_build       本地自动化测试，包括：init_env,vet_local,ut_coverage,package
    master_build      主干自动化测试，包括：init_env,vet_master,ut_coverage,package
    "
    return 0
}
# --------------------------------------------------

# --------------------------------------------------

case $1 in
    h|help|-h|-help)
        usage
    ;;
    buildd)
        exec_func "build"
    ;;
    buildr)
        exec_func "build"
    ;;
    ##如下命令，调用../uni-autotest/build_test/main.sh中的命令
    init_env)
        exec_command "init_env" 
    ;;
    init_test_env)
        exec_func "init_autotest_env"
    ;;
    package)
        exec_func "init_autotest_env"
        exec_command "package $2 $3" 
    ;;
    vet_local)
        exec_func "init_autotest_env"
        exec_command "vet_local" 
    ;;
    vet_master)
        exec_func "init_autotest_env"
        exec_command "vet_master" 
    ;;
    ut_coverage)
        exec_func "init_autotest_env"
        exec_command "ut_coverage" 
    ;;
    local_build)
        exec_func "init_autotest_env"
        exec_func "build"
        exec_command "local_build" 
    ;;
    master_build)
        exec_func "init_autotest_env"
        exec_func "build"
        exec_command "master_build $2 $3" 
    ;;
    tag_build)
        exec_func "init_autotest_env"
        exec_func "build"
        exec_command "master_build $2 $3" 
    ;;
    *)
        usage
    ;;
esac

exit 0
# --------------------------------------------------
