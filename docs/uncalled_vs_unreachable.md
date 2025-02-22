All uncalled functions are unreachable. (except for the default entrypoint)\
Not all unreachable functions are uncalled.

But by **iteratively removing uncalled** functions \
you should eventually reach a point where all unreachable \
functions have been removed.\
Because removing uncalled functions can potentially make \
other functions uncalled.

#### removing uncalled functions

![image info](./graphs/unreachable_1.svg)

![image info](./graphs/unreachable_2.svg)

#### libraries

In a library potentially every exported function can be called.\
So you will need to differentiate between main packages and libraries.
