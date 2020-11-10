import sys
import os
BOJ_NUM = sys.argv[1]

print(f'Begin testing {BOJ_NUM}....')

file_list = []
i : int = 1
while True:
    path = './{0}/{0}.in{1}'.format(BOJ_NUM, i)
    if os.path.isfile(path):
        file_list.append(path)
        i = i + 1
        # print(f'found file {path}')
    else:
        # print(f'not found file {path}')
        break

for test_file in file_list:
    exe_cmd = 'cat {1} | python3 ./{0}/{0}.py'.format(BOJ_NUM, test_file)
    os.system(exe_cmd)

print(f'End {BOJ_NUM}....')
