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
    gp.AddJob("test", func(ctx context.Context) error {
        select {
        case <-ctx.Done():
            log.Println("context closed")
            return nil
        default:
            panic("test panic")
            time.Sleep(time.Second * 10)
            return nil
        }
    })
    time.Sleep(time.Second)
    gp.ShutDown(true,time.Second) 

