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
				gp.AddJob("test", func(ctx context.Context, args ...interface{}) error {
					select {
					case <-ctx.Done():
						log.Println("shut down succeessfully")
						return nil
					default:
						time.Sleep(time.Second * 10)
						return nil
					}
				})
				gp.ShutDown(true, time.Second)

			})
		})

		Context("if we create an instance of gopool and create a job", func() {
			It("if the job panics we should exit safely", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())
				gp.AddJob("test", func(ctx context.Context, args ...interface{}) error {
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
				gp.ShutDown(true, time.Second)

			})
		})

		Context("if we create an instance of gopool and create multiple jobs", func() {
			It("if the job panics we should exit safely", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

				fn := func(ctx context.Context, args ...interface{}) error {
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
				gp.ShutDown(true, time.Second)
				duration := time.Since(before)

				Expect(duration).ShouldNot(BeNumerically(">", 5009171659))

			})
		})

		Context("if we create an instance of gopool and create multiple jobs", func() {
			It("if the job panics we should exit safely", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

				fn := func(ctx context.Context, args ...interface{}) error {
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
				gp.ShutDown(false, time.Second*2)
				duration := time.Since(before)

				Expect(duration).Should(BeNumerically("<", 2100059310))

			})
		})

		Context("if we create an instance of gopool and create multiple jobs with mutiple arguments", func() {
			It("then we should execute all of them sucessfully", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

				fn := func(ctx context.Context, args ...interface{}) error {
					select {
					case <-ctx.Done():
						time.Sleep(time.Second * 5)
						log.Println("the value passed is", "arg", args[2])

						log.Println("shut down succeessfully")
						return nil
					case <-args[0].(context.Context).Done():
						time.Sleep(time.Second * 5)

						log.Println("shut down succeessfully")
						return nil
					case <-args[1].(context.Context).Done():
						time.Sleep(time.Second * 5)

						log.Println("shut down succeessfully")
						return nil
					case <-time.After(time.Second * 10):
						panic("test panic")
						time.Sleep(time.Second * 10)
						return nil
					}
				}
				gp.AddJob("test", fn, context.Background(), context.Background(), 5)
				gp.AddJob("test1", fn, context.Background(), context.Background(), 5)
				gp.AddJob("test2", fn, context.Background(), context.Background(), 5)
				gp.ShutDown(true, time.Second)

			})
		})


		FContext("if we create an instance of gopool and create multiple jobs and call shutdown multiple times", func() {
			It("then we should not panic", func() {
				gp := NewGoPool(context.Background())
				Expect(gp).ShouldNot(BeNil())

				fn := func(ctx context.Context, args ...interface{}) error {
					select {
					case <-ctx.Done():
						time.Sleep(time.Second * 5)
						log.Println("the value passed is", "arg", args[2])

						log.Println("shut down succeessfully")
						return nil
					case <-args[0].(context.Context).Done():
						time.Sleep(time.Second * 5)

						log.Println("shut down succeessfully")
						return nil
					case <-args[1].(context.Context).Done():
						time.Sleep(time.Second * 5)

						log.Println("shut down succeessfully")
						return nil
					case <-time.After(time.Second * 10):
						panic("test panic")
						time.Sleep(time.Second * 10)
						return nil
					}
				}
				gp.AddJob("test", fn, context.Background(), context.Background(), 5)
				gp.AddJob("test1", fn, context.Background(), context.Background(), 5)
				gp.AddJob("test2", fn, context.Background(), context.Background(), 5)
				gp.ShutDown(true, time.Second)
				gp.ShutDown(true, time.Second)
				gp.ShutDown(true, time.Second)
				gp.ShutDown(true, time.Second)

			})
		})
	})

})
