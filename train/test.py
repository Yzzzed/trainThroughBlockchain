print("Hello World!")

'''
python distribute.py --ps_hosts=127.0.0.1:2222 --worker_hosts=127.0.0.1:2224,127.0.0.1:2225 --job_name=ps --task_index=0

python distribute.py --ps_hosts=127.0.0.1:2222 --worker_hosts=127.0.0.1:2224,127.0.0.1:2225 --job_name=worker --task_index=0

python distribute.py --ps_hosts=127.0.0.1:2222 --worker_hosts=127.0.0.1:2224,127.0.0.1:2225 --job_name=worker --task_index=1
'''