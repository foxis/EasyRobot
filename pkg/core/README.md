# Core of the Easy Robot framework

Features:
    [ ] Arbitrary value store
        [ ] (un)marshal to/from JSON
            [ ] Unittests
        [ ] (un)marshal to/from binary
            [ ] Unittests
    [ ] Logger wrapper
        [x] Empty Zerolog wrapper
        [x] Zerolog instance
    [ ] Transport wrapper
        [ ] TCP
        [ ] UDP
        [ ] Serial
        [ ] I2C
        [ ] SPI
    [ ] Plugin subsystem
        [x] Register plugins on import
            [ ] Unittests
        [x] Find and instantiate plugins
            [ ] Unittests
        [x] Plugin options marshalling
            [x] Unittests
    [ ] Pipeline subsystem
        [x] Step registry
            [ ] Unittests
        [x] Pipeline execution
            [ ] Unittests
        [ ] Pipeline configuration save/load
            [ ] Unittests
        [ ] Bridge using transport
            [ ] Unittests
        [ ] Steps
            [x] FPS counter
                [ ] Unittests
            [x] Synchronizer
                [ ] Unittests
            [x] Source
                [ ] Unittests
            [x] Sink
                [ ] Unittests
            [x] Fan In
                [ ] Unittests
            [x] Fan Out
                [ ] Unittests
            [x] Join
                [ ] Unittests
            [x] Processor
                [ ] Unittests
    [ ] Native Math Implementation (float32)
        [ ] Vector
        [ ] Matrix
        [ ] Tensor