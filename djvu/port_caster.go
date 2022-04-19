package djvu

import "sync"

// PortCaster maintains associations between ports.
//
// It monitors the status of all ports (have they been destructed yet?),
// accepts requests and notifications from them and forwards them to
// destinations according to internally maintained map of routes.
//
// The caller can modify the route map any way he likes
// (see AddRoute, DelRoute, CopyRoutes, etc. methods).
// Any port can be either a sender of a message,
// an intermediary receiver or a final destination.
//
// When a request is sent,
// the PortCaster computes the list of destinations by consulting with the route map.
// Notifications are only sent to ``alive'' ports.
// A port is alive if it is referenced by a valid pointer.
//  As a consequence, a port usually becomes alive after running the constructor
// (since the returned pointer is then assigned to a smartpointer)
// and is no longer alive when the port is destroyed
// (because it would not be destroyed if a smartpointer was referencing it).
//
// Destination ports are sorted according to their distance from the source.
// For example, if port `A` is connected to ports `B` and `C` directly,
// and port `B` is connected to `D`,
// then `B` and `C` are assumed to be one hop away from `A`,
// while `D` is two hops away from `A`.
//
// In some cases the requests and notifications are sent to every possible destination,
// and the order is not significant (like it is for NotifyFileFlagsChanged request).
// Others should be sent to the closest destinations first,
// and only then to the farthest, in case if they have not been processed by the closest.
// The examples are RequestData, NotifyError, and NotifyStatus.
//
// The user is not expected to create the PortCaster itself.
// He should use the global function `GetPortCaster()` instead.
type PortCaster struct {
	mtx      sync.Mutex
	routeMap map[Port][]Port

	// TODO: Do we even need this though?
	contMap map[Port]struct{}

	// Map of aliases to their respective ports
	a2pMap map[string]Port
}

// Global port caster
var globalPortCaster = &PortCaster{
	mtx:      sync.Mutex{},
	routeMap: make(map[Port][]Port),
	contMap:  make(map[Port]struct{}),
	a2pMap:   make(map[string]Port),
}

func GetPortCaster() *PortCaster {
	return globalPortCaster
}

// Removes the specified port from all routes.
// It will no longer be able to receive or generate messages
// and will be considered "dead" by `IsPortAlive`.
func (c *PortCaster) DelPort(p Port) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.clearAliases(p)
	delete(c.contMap, p)
}

// Adds route from `source` to `dest`.
// Whenever a request is sent or received by `source`,
// it will be forwarded to `dest` as well.
func (c *PortCaster) AddRoute(source Port, dest Port) {
	// Do things here...
	panic("unimplemented")
}

// The opposite of `AddRoute`.
// Removes the association between `source` and `dest`.
func (c *PortCaster) DelRoute(source Port, dest Port) {
	panic("unimplemented")
}

// Copies all incoming and outgoing routes from `source` to `dest`.
// This function should be called when a Port is copied,
// if you want to preserve the connectivity.
func (c *PortCaster) CopyRoutes(dest Port, source Port) {
	panic("unimplemented")
}

// Returns a smart pointer to the port if `p` is a valid pointer to an existing Port.
// Returns a null pointer otherwise.
//
// TODO: How do you translate this in a garbage-collected language like Go?
func (c *PortCaster) IsPortAlive(p Port) Port {
	panic("unimplemented")
}

// Assigns one more {\em alias} for the specified \Ref{DjVuPort}.
// {\em Aliases} are names, which can be used later to retrieve this
// \Ref{DjVuPort}, if it still exists. Any \Ref{DjVuPort} may have
// more than one {\em alias}. But every {\em alias} must correspond
// to only one \Ref{DjVuPort}. Thus, if the specified alias is
// already associated with another port, this association will be
// removed.
func (c *PortCaster) AddAlias(p Port, alias string) {
	panic("unimplemented")
}

// Removes all the aliases
func (c *PortCaster) ClearAllAliases() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.a2pMap = make(map[string]Port)
}

// Removes all aliases associated with the given Port
func (c *PortCaster) ClearAliases(p Port) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.clearAliases(p)
}

// Unsafe. Must only be called from DelPort or ClearAliases methods.
func (c *PortCaster) clearAliases(p Port) {
	for k, v := range c.a2pMap {
		if v == p {
			delete(c.a2pMap, k)
		}
	}
}

// Returns Port associated with the given alias.
// If nothing is known about name alias,
// or the port associated with it has already been destroyed,
// a nil pointer will be returned.
func (c *PortCaster) AliasToPort(alias string) Port {
	panic("unimplemented")
}

// Returns a list of Ports with aliases starting with `prefix`.
// If no Ports have been found, an empty list is returned.
func (c *PortCaster) PrefixToPorts(alias string) []Port {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest until one of them returns non-empty URL.
func (c *PortCaster) IdToUrl(source Port, id string) Url {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest until one of them returns non-zero pointer to a File.
func (c *PortCaster) IdToFile(source Port, id string) *File {
	panic("unimplemented")
}

// This request is issued when decoder needs additional data for decoding.
// Both File and Document are initialized with a URL, not the document data.
// As soon as they need the data, they call this function,
// whose responsibility is to locate the source of the data based on the URL passed
// and return it back in the form of the Pool.
// If this particular receiver is unable to fullfil the request, it should return nil.
func (c *PortCaster) RequestData(source Port, url *Url) *sync.Pool {
	panic("unimplemented")
}

// This notification is sent when an error occurs
// and the error message should be shown to the user.
// Returns whether the request is successful.
func (c *PortCaster) NotifyError(source Port, message string) bool {
	panic("unimplemented")
}

// This notification is sent to update the decoding status.
// Returns whether the request is successful.
func (c *PortCaster) NotifyStatus(source Port, message string) bool {
	panic("unimplemented")
}

// This notification is sent by an Image when it should be redrawn.
// It may be used to implement progressive redisplay.
func (c *PortCaster) NotifyRedisplay(source *Image) {
	panic("unimplemented")
}

// This notification is sent by Image
// when its geometry has been changed as a result of decoding.
// It may be used to implement progressive redisplay.
func (c *PortCaster) NotifyRelayout(source *Image) {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest.
func (c *PortCaster) NotifyChunkDone(source *Image, name string) {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest.
func (c *PortCaster) NotifyFileFlagsChanged(source *File, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest.
func (c *PortCaster) NotifyDocFlagsChanged(source *Document, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// Computes destination list for `source`
// and calls the corresponding function in each of the ports from the destination list
// starting from the closest.
func (c *PortCaster) NotifyDecodeProgress(source *Port, done float64) {
	panic("unimplemented")
}
