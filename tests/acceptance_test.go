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

	Describe("Check basic BOOK commands scenarios", func() {
		Context("BOOK A0 1 should result SUCCESS", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				session = executeCommands(cmd1Args)
			})

			It("exits with status code 0", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(0))
			})

			It("assert that ticket booking is successful", func() {
				expected := `SUCCESS`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A0 1, CANCEL A0 1, BOOK A0 1 should result in SUCCESS", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				cmd2Args := []string{execPath, "CANCEL", "A0", "1", "--state-file", tempDir}
				cmd2Arg := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				session = executeCommands(cmd1Args, cmd2Args, cmd2Arg)
			})

			It("exits with status code 0", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(0))
			})

			It("assert that ticket booking is successful", func() {
				expected := `SUCCESS`
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
				session = executeCommands(cmd1Args, cmd2Args)
			})

			It("exits with status code 1", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket booking fails", func() {
				expected := `FAIL`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A0 1, BOOK A1 1, BOOK A2 4 should result in SUCCESS", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A0", "1", "--state-file", tempDir}
				cmd2Args := []string{execPath, "BOOK", "A1", "1", "--state-file", tempDir}
				cmd3Args := []string{execPath, "BOOK", "A2", "4", "--state-file", tempDir}
				session = executeCommands(cmd1Args, cmd2Args, cmd3Args)
			})

			It("exits with status code 0", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(0))
			})

			It("assert that ticket booking is successful", func() {
				expected := `SUCCESS`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A2 4, BOOK A5 should result in FAIL", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A2", "4", "--state-file", tempDir}
				cmd2Args := []string{execPath, "BOOK", "A5", "1", "--state-file", tempDir}
				session = executeCommands(cmd1Args, cmd2Args)
			})

			It("exits with status code 1", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket booking fails", func() {
				expected := `FAIL`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A6 3 should result in FAIL", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A6", "3", "--state-file", tempDir}
				session = executeCommands(cmd1Args)
			})

			It("exits with status code 1", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket booking fails", func() {
				expected := `FAIL`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK A8 1 should result in FAIL", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "A8", "1", "--state-file", tempDir}
				session = executeCommands(cmd1Args)
			})

			It("exits with status code 1", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket booking fails", func() {
				expected := `FAIL`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
			})

			AfterEach(func() {
				gexec.CleanupBuildArtifacts()
				os.RemoveAll(tempDir)
			})
		})

		Context("BOOK U1 1 should result in FAIL", func() {
			var tempDir string
			var session *gexec.Session

			BeforeEach(func() {
				tempDir, err := os.MkdirTemp("", "booking-state")
				Expect(err).NotTo(HaveOccurred())
				cmd1Args := []string{execPath, "BOOK", "U1", "1", "--state-file", tempDir}
				session = executeCommands(cmd1Args)
			})

			It("exits with status code 1", func() {
				Eventually(session.Wait(10 * time.Second)).Should(gexec.Exit(1))
			})

			It("assert that ticket booking fails", func() {
				expected := `FAIL`
				Eventually(string(session.Wait(10 * time.Second).Out.Contents())).Should(ContainSubstring(expected))
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

func executeCommands(commands ...[]string) *gexec.Session {
	for i := 0; i < len(commands)-1; i++ {
		cmd1Args := commands[i]
		cmd1 := exec.Command(cmd1Args[0], cmd1Args[1:]...)

		session1, err := gexec.Start(cmd1, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		session1.Wait()
	}
	cmd2Args := commands[len(commands)-1]
	cmd2 := exec.Command(cmd2Args[0], cmd2Args[1:]...)

	session2, err := gexec.Start(cmd2, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session2
}
