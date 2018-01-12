package gotool

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

//本地执行命令
func GoCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

//密码验证方式登录远程服务器
func SSHPasswordLogin(hostAddr string, username string, password string) (*ssh.Session, error) {
	var session *ssh.Session
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: time.Second * 5,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	if client, err := ssh.Dial("tcp", hostAddr, config); err != nil {
		return session, err
	} else {
		return client.NewSession()
	}
}

//证书登录远程服务器
func SSHKeyLogin(hostAddr string, username string, keyFilePath string) (session *ssh.Session, err error) {
	key, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return session, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return session, err
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", hostAddr, config)
	if err != nil {
		return session, err
	}
	
	return client.NewSession()
}

//远程执行命令
func RemoteCommand(session *ssh.Session, command string) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	defer session.Close()
	if err := session.Run(command); err == nil {
		return b.String(), err
	} else {
		return "", err
	}
}

//获取当前运行路径
func GetWorkPath() string {
	if file, err := exec.LookPath(os.Args[0]); err == nil {
		return filepath.Dir(file) + "/"
	}
	return "./"
}
