package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGetReadmeRepoLengthHappyPath(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	url := "http://foo/bar/"
	// Content here is 75 lines
	data := `
	{
		"name": "README.md",
		"path": "README.md",
		"sha": "f6dc6f468503b63a5b8d69036a8d0dabe9fbaa33",
		"size": 2679,
		"url": "https://api.github.com/repos/spacemakerai/toadsworth/contents/README.md?ref=master",
		"html_url": "https://github.com/spacemakerai/toadsworth/blob/master/README.md",
		"git_url": "https://api.github.com/repos/spacemakerai/toadsworth/git/blobs/f6dc6f468503b63a5b8d69036a8d0dabe9fbaa33",
		"download_url": "https://raw.githubusercontent.com/spacemakerai/toadsworth/master/README.md?token=ADiGyFIC_roZh6u7CwgKQar0F7J3phHDks5bljmIwA%3D%3D",
		"type": "file",
		"content": "IyBUb2Fkc3dvcnRoCgpBZ2VudCBmb3IgY29sbGVjdGluZyBzeXN0ZW0vY29u\ndGFpbmVyIGFuZCBhcHBsaWNhdGlvbnMgbWV0cmljcyBmb3IgYSBub2RlLgpD\nb250YWluZXJzIGFuZCBpbnN0YW5jZXMgbm8gbG9uZ2VyIG5lZWQgdG8gZGVh\nbCB3aXRoIHNlcnZpY2UgZGlzY292ZXJ5LCBhcyBtZXRyaWNzCmNhbiBiZSBz\nZW50IGRpcmVjdGx5IHRvIFRvYWRzd29ydGguCgojIyBIb3cgaXQgd29ya3MK\nClRvYWRzd29ydGggaXMgYSBkZWFtb24gdGhhdCBydW5zIGFuIEhUVFAgc2Vy\ndmVyIG9uIHBvcnQgOTExMCAoY29uZmlndXJhYmxlKS4gVGhlCnNlcnZlciBj\nb2xsZWN0cyBib3RoIG1ldHJpY3MgZnJvbSB0aGUgbm9kZSAoQ1BVLCBtZW1v\ncnksIGV0Yy4pIGFuZCBmcm9tIHRoZQphcHBsaWNhdGlvbiBydW5uaW5nIG9u\nIHRoZSBub2RlLgoKIVtUb2Fkc3dvcnRoIERpYWdyYW1dKHRvYWRzd29ydGgu\ncG5nKQoKVGhlIG1ldHJpY3MgYXJlIGV4cG9zZWQgaW4gdGhlIGZvbGxvd2lu\nZyBlbmRwb2ludHM6CgoKfCBFbmRwb2ludCB8IERlc2NyaXB0aW9uIHwKfCAt\nLS0gfCAtLS0gfAp8IGBtZXRyaWNzL2FwcGxpY2F0aW9uLyBgIHwgTWV0cmlj\ncyBmcm9tIHRoZSBhcHBsaWNhdGlvbi4gVGhpcyBlbmRwb2ludCBpcyBhIHJl\ndmVyc2UgcHJveHkgdGhhdCB3aWxsIGZvcndhcmQgbWV0cmljcyBmcm9tIHRo\nZSBhcHBsaWNhdGlvbi4gfAp8IGBtZXRyaWNzL25vZGUvIGAgfCBNZXRyaWNz\nIGZyb20gdGhlIG5vZGUsIGxpa2UgQ1BVIGFuZCBtZW1vcnkgdXNhZ2UuICB8\nCgpJbiBhZGRpdGlvbiwgZHVlIHRvIGxpbWl0YXRpb25zIGluIHRoZSBuZXR3\nb3JrIGxheWVyIG9mIEFXUyBiYXRjaCwgVG9hZHN3b3J0aApzZW5kcyBhIGhl\nYXJ0YmVhdCBzaWduYWwgdG8gdGhlIFRvYWRzd29ydGggc2VydmljZSBkaXNj\nb3ZlcnkgZXZlcnkgMTAgc2Vjb25kcy4KSXQgd2lsbCBzZW5kIHRvIGAxMC4w\nLjIuODE6ODA4MGAsIHdoaWNoIGlzIHRoZSBQcm9tZXRoZXVzIGluc3RhbmNl\nIGhvc3RlZCBpbiBBV1MKVlBDIGB2cGMtMTI4ZGQzNzZgLiBUb2Fkc3dvcnRo\nIHdpbGwgY3VycmVudGx5IG5vdCB3b3JrIGluIGFueSBvdGhlciBlbnZpcm9u\nbWVudHMuIApUb2Fkc3dvcnRoIGFsc28gZmluZHMgYSBzdWl0YWJsZSBhbmQg\nb3BlbiBwb3J0LCB0aGlzIGFnYWluIGlzIHVlIHRvIGEgbmV0d29yayByZXN0\ncmljdGlvbiAKaW4gaG93IEFXUyBCYXRjaCBvcGVyYXRlcy4gVGhpcyBwb3J0\nIGlzIG5vdGlmaWVkIHRvIHRoZSBzZXJ2aWNlIGRpc2NvdmVyeS4KCgojIyBD\nb25maWcKClRvYWRzd29ydGggaXMgY29uZmlndXJlZCB0aHJvdWdoIHRoZSBm\nb2xsb3dpbmcgZW52aXJvbm1lbnQgdmFyaWFibGVzOgoKfCBFbnZpcm9ubWVu\ndCBWYXJpYWJsZSB8IERlZmF1bHQgdmFsdWUgfCBEZXNjcmlwdGlvbiB8Cnwg\nLS0tIHwgLS0tIHwgLS0tIHwKfCBgVE9BRFNXT1JUSF9BUFBMSUNBVElPTl9N\nRVRSSUNTX1VSTGAgfCBgaHR0cDovL2xvY2FsaG9zdDo5MTEwL21ldHJpY3Mv\nYCB8IFRoZSBVUkwgd2hlcmUgdG8gZmV0Y2ggdGhlIGFwcGxpY2F0aW9uIG1l\ndHJpY3MuIHwKCiMjIEFkZGluZyBUb2Fkc3dvcnRoIHRvIGRvY2tlcgpBZGQg\ndGhlIGJlbG93IGxpbmVzIHRvIHlvdXIgZG9ja2VyZmlsZSAoc3Vic3RpdHV0\nZSBge3ZlcnNpb259YCB3aXRoIHRoZSBsYXRlc3QKdmVyc2lvbiBmb3VuZCBp\nbiBbdGhlIGxhdGVzdCByZWxlYXNlXShodHRwczovL2dpdGh1Yi5jb20vc3Bh\nY2VtYWtlcmFpL3RvYWRzd29ydGgvcmVsZWFzZXMvbGF0ZXN0KToKYGBgCkFS\nRyBHSVRIVUJfVE9LRU4KRU5WIEdJVEhVQl9UT0tFTj0kR0lUSFVCX1RPS0VO\nIApSVU4gVkVSU0lPTj17dmVyc2lvbn0gY3VybCAtc0wgaHR0cHM6Ly8kR0lU\nSFVCX1RPS0VOQHJhdy5naXRodWJ1c2VyY29udGVudC5jb20vc3BhY2VtYWtl\ncmFpL3RvYWRzd29ydGgvbWFzdGVyL2luc3RhbGwuc2ggfCBzaApFWFBPU0Ug\nOTExMQpFTlRSWVBPSU5UIFsiL29wdC90b2Fkc3dvcnRoL2VudHJ5cG9pbnQu\nc2giXQpgYGAKClRoZSBhYm92ZSB3aWxsIHVzZSB0aGUgbGF0ZXN0IGluc3Rh\nbGwgc2NyaXB0IHRvIGluc3RhbGwgdGhlIHZlcnNpb24gc3BlY2lmaWVkIGlu\nCnRoZSBgVkVSU0lPTmAgZW52aXJvbm1lbnQgdmFyaWFibGUuCgpUbyBidWls\nZCB0aGUgZG9ja2VyIGltYWdlLCBwYXNzIHRoZSBgR0lUSFVCX1RPS0VOYCBh\ncyBhIGJ1aWxkIGFyZ3VtZW50OgpgYGAKZG9ja2VyIGJ1aWxkIC0tYnVpbGQt\nYXJnIEdJVEhVQl9UT0tFTj0kR0lUSFVCX1RPS0VOLi4uCmBgYAoKIyMgUmVs\nZWFzaW5nClRvYWRzd29ydGggd2lsbCBiZSByZWxlYXNlZCBvbiBldmVyeSBt\nZXJnZSB0byBtYXN0ZXIuIFRoZSB2ZXJzaW9uIHdpbGwgYmUgb24gdGhlCmZv\ncm1hdCBgTUFKT1IuTUlOT1IuVElNRVNUQU1QYC4gYE1BSk9SYCBhbmQgYE1J\nTk9SYCBpcyBkZWZpbmVkIGluIHRoZSBgVkVSU0lPTmAKZmlsZSBpbiB0aGUg\ncm9vdCBvZiB0aGlzIHJlcG9zaXRvcnkuIGBUSU1FU1RBTVBgIGlzIGRlZmlu\nZWQgZGVwbG95IHRpbWUgYnkKdGhlIENJIHNlcnZpY2UgdXNpbmcgYGRhdGUg\nKyVzYC4KCiMjIEhhY2tpbmcKSWYgeW91J3JlIGludGVyZXN0ZWQgaW4gaGFj\na2luZyBvbiBUb2Fkc3dvcnRoLCBpbnN0YWxsIHRoZSBkZXBlbmRlbmNpZXMg\nd2l0aApgZGVwYDoKCmBgYGJhc2gKJCBkZXAgZW5zdXJlCmBgYAoKUnVuOgpg\nYGBiYXNoCiQgLi9zdGFydC1kZXYKYGBg\n",
		"encoding": "base64",
		"_links": {
		  "self": "https://api.github.com/repos/spacemakerai/toadsworth/contents/README.md?ref=master",
		  "git": "https://api.github.com/repos/spacemakerai/toadsworth/git/blobs/f6dc6f468503b63a5b8d69036a8d0dabe9fbaa33",
		  "html": "https://github.com/spacemakerai/toadsworth/blob/master/README.md"
		}
	  }
 	`
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, data))
	res, err := getRepoReadmeLength(url)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 75, res)
}
