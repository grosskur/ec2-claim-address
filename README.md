# ec2-claim-address

Given a list of [Elastic IP][elastic-ip] addresses, try to associate
one with the current [EC2][ec2] instance.

## Getting started

### Precompiled binaries

* [Linux (64-bit)](https://github.com/grosskur/ec2-claim-address/releases/download/v20140307/ec2-claim-address)

### Compile from source

```bash
$ go get -u github.com/grosskur/ec2-claim-address
```

### Usage

```bash
$ ec2-claim-address ADDRESS1 [ADDRESS2 ...]
```

### Example

```bash
$ ec2-claim-address 54.85.21.92 54.84.36.41
ec2-claim-address: getting instance-id from metadata
ec2-claim-address: found instance-id: i-9e2790df
ec2-claim-address: connecting to ec2 API endpoint
ec2-claim-address: getting address information
ec2-claim-address: claiming unassociated address: 54.84.36.41
ec2-claim-address: successfully associated: 54.84.36.41
```

## Background

An [Elastic IP][elastic-ip] is a static IP that can be remapped from
one instance to another. However, an Elastic IP can only be assigned
*after* an instance is created. Suppose you have an [Auto Scaling
Group][auto-scaling-group] and want to assign an Elastic IP to each
instance. You can run `ec2-claim-address` with a list of possible
Elastic IPs from the [User Data][user-data] script to automatically
pick an unassociated one.

[auto-scaling-group]: http://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/WorkingWithASG.html
[ec2]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html
[elastic-ip]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html
[iam-role]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html
[instance-metadata]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html
[user-data]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html
