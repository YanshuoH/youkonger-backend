#!/usr/bin/env python2

from fabric.api import task, local

@task
def test(cover=''):
    local("rm -f coverage/html/*.html")
    local("rm -f coverage/*.out")

    pkgs = {
        'youkonger-conf': "github.com/YanshuoH/youkonger/conf",
        'youkonger-dao': "github.com/YanshuoH/youkonger/dao",
        'youkonger-consts': "github.com/YanshuoH/youkonger/consts",
        # 'youkonger-controllers': "github.com/YanshuoH/youkonger/controllers",
        'youkonger-forms': "github.com/YanshuoH/youkonger/forms",
        'youkonger-utils': "github.com/YanshuoH/youkonger/utils",
    }

    for key, pkg in pkgs.iteritems():
        if cover == '':
            local("go test %s" % (pkg))
            continue
        local("go test -coverprofile=coverage/%s.out %s" % (key, pkg))
        local("go tool cover -html=coverage/%s.out -o coverage/html/%s.html" % (key, key))