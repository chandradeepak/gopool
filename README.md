# gopool

It allows to create go routines which are panic 
tolerant. we can specify a panic handler for the
panics to be handled.

All the go routines that are created by gopool 
to exit cleanly by issuing a shutdown.

once a shut down is issued it waits for all go 
routines to exit before it can come out or it waits
till the timeout mentioned elapsed.



usage
-----
    gp := NewGoPool(context.Background())
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

