# rbgtest

Simple Go package to test Jenkins and Bamboo jobs with tests using go2xunit. <code>RandomBit()</code> randomly returns 0 or 1. Test fails if the returned value is 0.

Add the following to a test job:

<code>go test -v | tee gotest.out | go2xunit -bamboo=true -output tests.xml</code>

# Prerequisites

<code>go get bitbucket.org/tebeka/go2xunit</code>

and update PATH:

<code>export PATH=$GOPATH/bin:$PATH</code>
