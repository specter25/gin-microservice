## Deployment instruction 

### Platfrom specific deployment (AWS EBS) without using the docker file
    - create a ebs policy , group(attach the policy to this user group) and attach a user to that group to give him the required permissions .
    - run on port 5000 (default port forwarding ebs)
    - make the build.sh shell script
    - make a Builfile to specify which .sh file EBS should run to build the project
    - make a Profile to specify the command to run the server

        - $ eb init 
            Then select the region and environments
            - rest defaults 
            - code-commit no 
            - ssh yes
        - $ eb create --single 
            - spot fleet No   

### Docker Deployment
    