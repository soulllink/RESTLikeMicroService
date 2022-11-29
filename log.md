https://kenyaappexperts.com/blog/how-to-deploy-golang-to-production-step-by-step/

excelebi@s2e24dfb0:/home$ systemctl stop dqmicro
Failed to stop dqmicro.service: The name org.freedesktop.PolicyKit1 was not provided by any .service files
See system logs and 'systemctl status dqmicro.service' for details.
excelebi@s2e24dfb0:/home$ sudo systemctl stop dqmicro
excelebi@s2e24dfb0:/home$ sudo systemctl disable dqmicro
Removed /etc/systemd/system/multi-user.target.wants/dqmicro.service.
excelebi@s2e24dfb0:/home$  


*****************
Description= instance to serve jobs api
After=network.target

[Service]
User=root
Group=www-data
ExecStart=/home/path/to/binary/you/uploaded(which in my case is/var/www/go/jobs)
[Install]
WantedBy=multi-user.target
*******************
Description=Discord Ladushki Bot instance
After=network.target

[Service]



User=excelebi
Group=www-data
ExecStart=/home/excelebi/work/bin/Stb
WorkingDirectory=imagefolderpath
[Install]
WantedBy=multi-user.target
**********************
excelebi@s2e24dfb0:~/work/src/DqMicroS$ sudo nano /etc/systemd/system/dqmicro.service
[sudo] password for excelebi:
excelebi@s2e24dfb0:~/work/src/DqMicroS$ sudo systemctl start dqmicro
excelebi@s2e24dfb0:~/work/src/DqMicroS$ sudo systemctl enable dqmicro
Created symlink /etc/systemd/system/multi-user.target.wants/dqmicro.service → /etc/systemd/system/dqmicro.service.
excelebi@s2e24dfb0:~/work/src/DqMicroS$   