package sio

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(s socketio.Conn) error {
	s.SetContext("")
	fmt.Println("connected:", s.ID())
	return nil
}

func OnDisconnect(s socketio.Conn, msg string) {
	fmt.Println("disconnect {}", s.ID())
	return
}

//@sio.on('authorize', namespace='/host')
func AuthorizeEvent(s socketio.Conn, msg string) error {
	fmt.Println("authorize ", s.ID())
	//users[sid] = {}
	//users[sid]['authorize'] = msg
	fmt.Println("authorize: {}", msg)
	//sio.emit('authorize', data={'msg': 'ok!'}, room=sid, namespace='/host')
	/*
			# read,{system:'vmstats'}
		    # sio.emit('read', {'system':'vmstats'}, room=sid, namespace='/host')
		    # sio.emit('run',{'cmd_seq_id':123,'act':'system','mtd':'reboot', 'form':{}})
		    # pass
	*/
	return nil
}

/*
# sio.emit('read',{system:'vmstats'})
@sio.on('resp', namespace='/host')
def resp_event(sid, data):
    """
    自定义事件消息的处理方法
    :param sid: string sid是发送此事件消息的客户端id
    :param data: data是客户端发送的消息数据
    """
    users[sid]['resp'] = data
    print('resp: {}'.format(json2txt(data)))
    # sio.emit('AUTH',{})
    #pass
    return


@sio.on('act', namespace='/host')
def act_event(sid, data):
    """
    自定义事件消息的处理方法
    :param sid: string sid是发送此事件消息的客户端id
    :param data: data是客户端发送的消息数据
    """
    users[sid]['act'] = data
    print('act: {}'.format(json2txt(data)))
    # sio.emit('AUTH',{})
    #pass
    return


@sio.on('mod', namespace='/host')
def mod_event(sid, data):
    """
    自定义事件消息的处理方法
    :param sid: string sid是发送此事件消息的客户端id
    :param data: data是客户端发送的消息数据
    """
    users[sid]['mod'] = data
    print('mod: {}'.format(json2txt(data)))
    # sio.emit('AUTH',{})
    #pass
    return

*/
//@sio.on('commit', namespace='/host')
func CommitEvent(s socketio.Conn, msg string) error {
	//users[sid]['commit'] = data
	fmt.Println("commit:  ", msg)
	//netconf = data.get('netconf')
	//if netconf:
	//    mac = netconf.get('mac')
	//else:
	//    return
	//if data['svrconn'].get('svrsync') == "syncok":
	//    imgupd = machines[mac].get('imgupd')
	/*
	   #if imgupd == 'upload':
	   #    # 上传完成,合并镜像，合并完成修改状态off
	   #    imgupdate.run("w10x64v1809bios")
	   #    machines[mac]['imgupd']='off'
	*/
	/*
			# sio.emit('AUTH',{})
		    # sio.emit('read', {'system':'vmstats'}, room=sid, namespace='/host')
		    # sio.emit('mod', {'act': 'system', 'mtd': 'gets', 'form': {'names': 'vmstats'}}, room=sid, namespace='/host')
		    # sio.emit('mod', {'act': 'system', 'mtd': 'gets', 'form': {'names': 'vmstats'}}, room=sid, namespace='/host')
		    # pass*/
	return nil

}
