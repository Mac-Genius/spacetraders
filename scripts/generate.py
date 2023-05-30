# Create generate.go under the api directory.
import os

def main():
    api_dir = os.path.join(os.path.dirname(os.path.dirname(__file__)), 'api')
    protos = [x for x in os.listdir(os.path.join(api_dir, 'proto')) if x.endswith('.proto')]
    protos.sort()
    with open(os.path.join(api_dir, 'generate.go'), 'w') as f:
        f.write('package api\n\n')
        for x in protos:
            f.write(f'//go:generate protoc -I=proto --go_out=./ --go_opt=paths=source_relative proto/{x}\n')

if __name__ == '__main__':
    main()
