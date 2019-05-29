import socket
import os

TCP_PORT = 5900
CURRENT_INDEX = 0
TIMEOUT = 3

def checkSecurityType(TCP_IP):
    #types 0 Invalid | 1 NONE | 2 VNC Authentication
    authentication = 0
    vnc = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    vnc.settimeout(TIMEOUT)
    try:
        vnc.connect( ( TCP_IP, int( TCP_PORT ) ) )
        vnc_ver = vnc.recv(12)
        print(vnc_ver)
        if 'RFB' not in str(vnc_ver):
            return authentication
        vnc.send(vnc_ver)
        types = ord( vnc.recv(1) )
        if not types:
            return authentication
        # print("security types: " + str( types ))
        for i in range(0,types):
            if ord( vnc.recv(1) ) == 1:
                authentication = 1
            else:
                pass
        vnc.shutdown(socket.SHUT_WR)
        vnc.close()
    except socket.error:
        vnc.close()
        pass
    except socket.timeout:
        vnc.close()
        pass
    return authentication

if __name__ == '__main__':
    try: input = raw_input
    except NameError: pass
    file_path = input("Please enter full file path to list of IPs.\nExample: /home/user/Documents/list.txt\nFull path: ")
    print("Note ... current index will be kept in new file 'vnc-log.txt' if for some reason the program is interrupted")
    print("\nRunning ...")
    with open(file_path, 'r') as file:
        for line in file:
            line_ip = line.strip('\n')
            vncType = checkSecurityType(line_ip)
            CURRENT_INDEX = CURRENT_INDEX + 1
            os.system("echo " + str(CURRENT_INDEX) + " > vnc-log.txt")
            if vncType == 1:
                command = "timeout 60 vncsnapshot -allowblank -vncQuality 6 -quality 60 " + line_ip + " " + line_ip + ".jpg> /dev/null 2>&1"
                os.system(command)
                print("done: " + line_ip)
            else:
                pass
    print("Program complete, enjoy sifting through your results :-)")
