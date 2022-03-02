package djvu

/** Thread safe data storage.
    The purpose of #DataPool# is to provide a uniform interface for
    accessing data from decoding routines running in a multi-threaded
    environment. Depending on the mode of operation it may contain the
    actual data, may be connected to another #DataPool# or may be mapped
    to a file. Regardless of the mode, the class returns data in a
    thread-safe way, blocking reading threads if there is no data of
    interest available. This blocking is especially useful in the
    networking environment (plugin) when there is a running decoding thread,
    which wants to start decoding as soon as there is just one byte available
    blocking if necessary.

    Access to data in a #DataPool# may be direct (Using \Ref{get_data}()
    function) or sequential (See \Ref{get_stream}() function).

    If the #DataPool# is not connected to anything, that is it contains
    some real data, this data can be added to it by means of two
    \Ref{add_data}() functions. One of them adds data sequentially maintaining
    the offset of the last block of data added by it. The other can store
    data anywhere. Thus it's important to realize, that there may be "white
    spots" in the data storage.

    There is also a way to test if data is available for some given data
    range (See \Ref{has_data}()). In addition to this mechanism, there are
    so-called {\em trigger callbacks}, which are called, when there is
    all data available for a given data range.

    Let us consider all modes of operation in details:

    \begin{enumerate}
       \item {\bf Not connected #DataPool#}. In this mode the #DataPool#
             contains some real data. As mentioned above, it may be added
             by means of two functions \Ref{add_data}() operating independent
	     of each other and allowing to add data sequentially and
	     directly to any place of data storage. It's important to call
	     function \Ref{set_eof}() after all data has been added.

	     Functions like \Ref{get_data}() or \Ref{get_stream}() can
	     be used to obtain direct or sequential access to the data. As
	     long as \Ref{is_eof}() is #FALSE#, #DataPool# will block every
	     reader, which is trying to read unavailable data until it
	     really becomes available. But as soon as \Ref{is_eof}() is
	     #TRUE#, any attempt to read non-existing data will read #0# bytes.

	     Taking into account the fact, that #DataPool# was designed to
	     store DjVu files, which are in IFF formats, it becomes possible
	     to predict the size of the #DataPool# as soon as the first
	     #32# bytes have been added. This is invaluable for estimating
	     download progress. See function \Ref{get_length}() for details.
	     If this estimate fails (which means, that stored data is not
	     in IFF format), \Ref{get_length}() returns #-1#.

	     Triggers may be added and removed by means of \Ref{add_trigger}()
	     and \Ref{del_trigger}() functions. \Ref{add_trigger}() takes
	     a data range. As soon as all data in that data range is
	     available, the trigger callback will be called.

	     All trigger callbacks will be called when #EOF# condition
	     has been set.

       \item {\bf #DataPool# connected to another #DataPool#}. In this
             {\em slave} mode you can map a given #DataPool# to any offsets
	     range inside another #DataPool#. You can connect the slave
	     #DataPool# even if there is no data in the master #DataPool#.
	     Any \Ref{get_data}() request will be forwarded to the master
	     #DataPool#, and it will be responsible for blocking readers
	     trying to access unavailable data.

	     The usage of \Ref{add_data}() functions is prohibited for
	     connected #DataPool#s.

	     The offsets range used to map a slave #DataPool# can be fully
	     specified (both start offset and length are positive numbers)
	     or partially specified (the length is negative). In this mode
	     the slave #DataPool# is assumed to extend up to the end
	     of the master #DataPool#.

	     Triggers may be used with slave #DataPool#s as well as with
	     the master ones.

	     Calling \Ref{stop}() function of a slave will stop only the slave
	     (and any other slave connected to it), but not the master.

	     \Ref{set_eof}() function is meaningless for slaves. They obtain
	     the #ByteStream::EndOfFile# status from their master.

	     Depending on the offsets range passed to the constructor,
	     \Ref{get_length}() returns different values. If the length
	     passed to the constructor was positive, then it is returned
	     by \Ref{get_length}() all the time. Otherwise the value returned
	     is either #-1# if master's length is still unknown (it didn't
	     manage to parse IFF data yet) or it is calculated as
	     #masters_length-slave_start#.

       \item {\bf #DataPool# connected to a file}. This mode is quite similar
             to the case, when the #DataPool# is connected to another
	     #DataPool#. Similarly, the #DataPool# stores no data inside.
	     It just forwards all \Ref{get_data}() requests to the underlying
	     source (a file in this case). Thus these requests will never
	     block the reader. But they may return #0# if there is no data
	     available at the requested offset.

	     The usage of \Ref{add_data}() functions is meaningless and
	     is prohibited.

	     \Ref{is_eof}() function always returns #TRUE#. Thus \Ref{set_eof}()
	     us meaningless and does nothing.

	     \Ref{get_length}() function always returns the file size.

	     Calling \Ref{stop}() function will stop this #DataPool# and
	     any other slave connected to it.

	     Trigger callbacks passed through \Ref{add_trigger}() function
	     are called immediately.

	     This mode is useful to read and decode DjVu files without reading
	     and storing them in full in memory.
    \end{enumerate}
*/
type DataPool struct {
	// TODO: The needful
}
