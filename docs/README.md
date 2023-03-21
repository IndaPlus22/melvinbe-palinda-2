## Task 2 - Many Senders; Many Receivers

The program many2many.go contains four producers that together send 32 strings over a channel. At the other end there are two consumers that receive the strings. Describe what happens, and explain why it happens, if you make the following changes in the program. Try first to reason your way through, and then test your hypothesis by changing and running the program.

* What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?
    * The channel will close before the producers have sent their information. The program will then panic as producers try to send on the closed channel.

* What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?
    * The first producer to finish will close the channel which will cause the other producers to send to a closed channel and the program will panic. 

* What happens if you remove the statement `close(ch)` completely?
    * There will not be much of a difference as the wg.Wait() call makes sure that all producers send before the channel is closed and so the program will not panic. The program also exits right after so it will not make much of a difference if the channel remains open or not.

* What happens if you increase the number of consumers from 2 to 4?
    * The program should still work and should even be faster as there would be less waiting for consumers to finish.

* Can you be sure that all strings are printed before the program
  stops?
    * No. The program only waits for the producers to finish sending but not for the consumers to print. This can be fixed by adding a second waitgroup for the consumers. 