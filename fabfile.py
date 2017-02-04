#!/usr/bin/env python2

from fabric.api import task, local, sudo, put, env

env.use_ssh_config = True

BIN_FILENAME = 'youkonger-bin'
CONF_FILENAME = 'conf_prod.toml'
TASK_NAME = 'youkonger'

@task
def build():
    local('GOOS=linux GOARCH=amd64 go build -o %s' % BIN_FILENAME)

@task
def clean():
    local('rm -f %s && go clean' % BIN_FILENAME)

@task
def deploy():
    clean()
    build()
    sudo('supervisorctl stop %s' % TASK_NAME)
    put(BIN_FILENAME, '~/youkonger/%s' % BIN_FILENAME)
    put('conf/%s' % CONF_FILENAME, '~/youkonger/%s' % CONF_FILENAME)
    put('views', '~/youkonger/')
    sudo('rm -f ~/youkonger/views/index.html')
    sudo('ln -s ~/youkonger-fe/dist/index.html ~/youkonger/views/index.html')
    sudo('rm -f ~/youkonger/public/assets')
    sudo('ln -s ~/youkonger-fe/dist ~/youkonger/public/assets')
    sudo('supervisorctl start %s' % TASK_NAME)
    clean()

@task
def test(cover=''):
    local("rm -f coverage/html/*.html")
    local("rm -f coverage/*.out")

    pkgs = {
        'youkonger-conf': "github.com/YanshuoH/youkonger/conf",
        'youkonger-dao': "github.com/YanshuoH/youkonger/dao",
        'youkonger-consts': "github.com/YanshuoH/youkonger/consts",
        'youkonger-api': "github.com/YanshuoH/youkonger/controllers/api",
        'youkonger-middlewares': "github.com/YanshuoH/youkonger/controllers/middlewares",
        'youkonger-forms': "github.com/YanshuoH/youkonger/forms",
        'youkonger-utils': "github.com/YanshuoH/youkonger/utils",
        'youkonger-jrender': "github.com/YanshuoH/youkonger/jrender",
    }

    for key, pkg in pkgs.iteritems():
        if cover == '':
            local("go test %s" % (pkg))
            continue
        local("go test -coverprofile=coverage/%s.out %s" % (key, pkg))
        local("go tool cover -html=coverage/%s.out -o coverage/html/%s.html" % (key, key))