package djvu

// Port provides base functionality for classes willing to take part in
// sending and receiving messages generated during decoding process.  You
// need to derive your class from Port if you want it to be able to
// send or receive requests. In addition, for receiving requests you should
// override one or more virtual function.
//
// Note: All ports should be allocated on the heap using
// #operator new# and immediately secured using a \Ref{GP} smart pointer.
// Ports which are not secured by a smart-pointer are not considered
// ``alive'' and never receive notifications!
type Port interface {
}

// TODO: How on earth do you even implement the destructor?
