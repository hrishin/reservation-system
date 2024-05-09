package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os"
	"os/exec"
	"time"
)

var _ = Describe("Reservation system tests", func() {
	var execPath string
	BeforeEach(func() {
		execPath = buildBinary()
	})

	Describe("Check BOOK commands", func() {
		Context("BOOK A0 1 should result SUCCESS", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				session = executeBinary(execPath, "BOOK", "A0", "1", "--state-file", tempDir)
			})

			It("exits with status code 0", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(0))
			})

			It("assert that ticket is booked successfully", func() {
				expected := `confirmed 1 tickets for seating A0`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A0 1 twice should result in FAIL", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				cmd2Args := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				session = executeBinaryTwoCommands(cmd1Args, cmd2Args)
			})

			It("exits with status code 0", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket is booked successfully", func() {
				expected := `booking failed for 1 tickets for seat A0 : requested seat number request is already booked: A0
ERR: requested seat number request is already booked: A0`
				Eventually(string(session.Wait(10 * time.Second).Err.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})
	})

	AfterEach(func() {
		os.RemoveAll(execPath)
	})
})

func buildBinary() string {
	dockerfileSourcesPath, err := gexec.Build("../main.go")
	Expect(err).NotTo(HaveOccurred())

	return dockerfileSourcesPath
}

func executeBinary(path string, args ...string) *gexec.Session {
	cmd := exec.Command(path, args...)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func executeBinaryTwoCommands(cmd1Args []string, cmd2Args []string) *gexec.Session {
	cmd1 := exec.Command(cmd1Args[0], cmd1Args[1:]...)
	cmd2 := exec.Command(cmd2Args[0], cmd2Args[1:]...)

	session1, err := gexec.Start(cmd1, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	session1.Wait()

	session2, err := gexec.Start(cmd2, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session2
}
