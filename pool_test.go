package gopool

import (
	"context"
	"log"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoPool Test", func() {

	Describe("Given we can create an instance of gopool", func() {
		Context("if we create an instance of gopool", func() {
			It("should not be nil", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

			})
		})

		Context("if we create an instance of gopool", func() {
			It("we should be able add new jobs", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())
				gp.AddJob("test", func(ctx context.Context) error {
					select {
					case <-ctx.Done():
						log.Println("shut down succeessfully")
						return nil
					default:
						time.Sleep(time.Second * 10)
						return nil
					}
				})
				gp.ShutDown()

			})
		})

		Context("if we create an instance of gopool and create a job", func() {
			It("if the job panics we should exit safely", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())
				gp.AddJob("test", func(ctx context.Context) error {
					select {
					case <-ctx.Done():
						log.Println("shut down succeessfully")
						return nil
					default:
						panic("test panic")
						time.Sleep(time.Second * 10)
						return nil
					}
				})
				time.Sleep(time.Second)
				gp.ShutDown()

			})
		})

		Context("if we create an instance of gopool and create multiple jobs", func() {
			It("if the job panics we should exit safely", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

				fn := func(ctx context.Context) error {
					select {
					case <-ctx.Done():
						time.Sleep(time.Second * 5)

						log.Println("shut down succeessfully")
						return nil
					case <-time.After(time.Second * 10):
						panic("test panic")
						time.Sleep(time.Second * 10)
						return nil
					}
				}
				gp.AddJob("test", fn)
				gp.AddJob("test1", fn)
				gp.AddJob("test2", fn)
				before := time.Now()
				gp.ShutDown()
				duration := time.Since(before)

				Expect(duration).ShouldNot(BeNumerically(">", 5009171659))

			})
		})
	})

})
