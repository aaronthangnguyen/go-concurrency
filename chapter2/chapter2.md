# CHAPTER 2

- New process copies memory, file handles, stack space, registers, and program counter
- New thread shares memory and file handles, but has own stack spacem registers, and program counter.
- Stack space stores local variables living within a function. Threads do not share stack space; though, they share memory space.
