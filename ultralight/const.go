// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 07 Oct 2019 13:59:36 CDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ultralight

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L${SRCDIR}/libs -lUltralightCore -lWebCore -lUltralight -lAppCore
#include "AppCore/CAPI.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// JSC_OBJC_API_ENABLED as defined in JavaScriptCore/JSBase.h:141
	JSC_OBJC_API_ENABLED = 0
)

// ULWindowFlags as declared in AppCore/CAPI.h:47
type ULWindowFlags int32

// ULWindowFlags enumeration from AppCore/CAPI.h:47
const (
	KWindowFlags_Borderless  ULWindowFlags = 1
	KWindowFlags_Titled      ULWindowFlags = 2
	KWindowFlags_Resizable   ULWindowFlags = 4
	KWindowFlags_Maximizable ULWindowFlags = 8
)

// ULMessageSource as declared in Ultralight/CAPI.h:73
type ULMessageSource int32

// ULMessageSource enumeration from Ultralight/CAPI.h:73
const (
	KMessageSource_XML            ULMessageSource = iota
	KMessageSource_JS             ULMessageSource = 1
	KMessageSource_Network        ULMessageSource = 2
	KMessageSource_ConsoleAPI     ULMessageSource = 3
	KMessageSource_Storage        ULMessageSource = 4
	KMessageSource_AppCache       ULMessageSource = 5
	KMessageSource_Rendering      ULMessageSource = 6
	KMessageSource_CSS            ULMessageSource = 7
	KMessageSource_Security       ULMessageSource = 8
	KMessageSource_ContentBlocker ULMessageSource = 9
	KMessageSource_Other          ULMessageSource = 10
)

// ULMessageLevel as declared in Ultralight/CAPI.h:81
type ULMessageLevel int32

// ULMessageLevel enumeration from Ultralight/CAPI.h:81
const (
	KMessageLevel_Log     ULMessageLevel = 1
	KMessageLevel_Warning ULMessageLevel = 2
	KMessageLevel_Error   ULMessageLevel = 3
	KMessageLevel_Debug   ULMessageLevel = 4
	KMessageLevel_Info    ULMessageLevel = 5
)

// ULCursor as declared in Ultralight/CAPI.h:128
type ULCursor int32

// ULCursor enumeration from Ultralight/CAPI.h:128
const (
	KCursor_Pointer                  ULCursor = iota
	KCursor_Cross                    ULCursor = 1
	KCursor_Hand                     ULCursor = 2
	KCursor_IBeam                    ULCursor = 3
	KCursor_Wait                     ULCursor = 4
	KCursor_Help                     ULCursor = 5
	KCursor_EastResize               ULCursor = 6
	KCursor_NorthResize              ULCursor = 7
	KCursor_NorthEastResize          ULCursor = 8
	KCursor_NorthWestResize          ULCursor = 9
	KCursor_SouthResize              ULCursor = 10
	KCursor_SouthEastResize          ULCursor = 11
	KCursor_SouthWestResize          ULCursor = 12
	KCursor_WestResize               ULCursor = 13
	KCursor_NorthSouthResize         ULCursor = 14
	KCursor_EastWestResize           ULCursor = 15
	KCursor_NorthEastSouthWestResize ULCursor = 16
	KCursor_NorthWestSouthEastResize ULCursor = 17
	KCursor_ColumnResize             ULCursor = 18
	KCursor_RowResize                ULCursor = 19
	KCursor_MiddlePanning            ULCursor = 20
	KCursor_EastPanning              ULCursor = 21
	KCursor_NorthPanning             ULCursor = 22
	KCursor_NorthEastPanning         ULCursor = 23
	KCursor_NorthWestPanning         ULCursor = 24
	KCursor_SouthPanning             ULCursor = 25
	KCursor_SouthEastPanning         ULCursor = 26
	KCursor_SouthWestPanning         ULCursor = 27
	KCursor_WestPanning              ULCursor = 28
	KCursor_Move                     ULCursor = 29
	KCursor_VerticalText             ULCursor = 30
	KCursor_Cell                     ULCursor = 31
	KCursor_ContextMenu              ULCursor = 32
	KCursor_Alias                    ULCursor = 33
	KCursor_Progress                 ULCursor = 34
	KCursor_NoDrop                   ULCursor = 35
	KCursor_Copy                     ULCursor = 36
	KCursor_None                     ULCursor = 37
	KCursor_NotAllowed               ULCursor = 38
	KCursor_ZoomIn                   ULCursor = 39
	KCursor_ZoomOut                  ULCursor = 40
	KCursor_Grab                     ULCursor = 41
	KCursor_Grabbing                 ULCursor = 42
	KCursor_Custom                   ULCursor = 43
)

// ULBitmapFormat as declared in Ultralight/CAPI.h:133
type ULBitmapFormat int32

// ULBitmapFormat enumeration from Ultralight/CAPI.h:133
const (
	KBitmapFormat_A8    ULBitmapFormat = iota
	KBitmapFormat_RGBA8 ULBitmapFormat = 1
)

// ULKeyEventType as declared in Ultralight/CAPI.h:140
type ULKeyEventType int32

// ULKeyEventType enumeration from Ultralight/CAPI.h:140
const (
	KKeyEventType_KeyDown    ULKeyEventType = iota
	KKeyEventType_KeyUp      ULKeyEventType = 1
	KKeyEventType_RawKeyDown ULKeyEventType = 2
	KKeyEventType_Char       ULKeyEventType = 3
)

// ULMouseEventType as declared in Ultralight/CAPI.h:146
type ULMouseEventType int32

// ULMouseEventType enumeration from Ultralight/CAPI.h:146
const (
	KMouseEventType_MouseMoved ULMouseEventType = iota
	KMouseEventType_MouseDown  ULMouseEventType = 1
	KMouseEventType_MouseUp    ULMouseEventType = 2
)

// ULMouseButton as declared in Ultralight/CAPI.h:153
type ULMouseButton int32

// ULMouseButton enumeration from Ultralight/CAPI.h:153
const (
	KMouseButton_None   ULMouseButton = iota
	KMouseButton_Left   ULMouseButton = 1
	KMouseButton_Middle ULMouseButton = 2
	KMouseButton_Right  ULMouseButton = 3
)

// ULScrollEventType as declared in Ultralight/CAPI.h:158
type ULScrollEventType int32

// ULScrollEventType enumeration from Ultralight/CAPI.h:158
const (
	KScrollEventType_ScrollByPixel ULScrollEventType = iota
	KScrollEventType_ScrollByPage  ULScrollEventType = 1
)

// JSType as declared in JavaScriptCore/JSValueRef.h:53
type JSType int32

// JSType enumeration from JavaScriptCore/JSValueRef.h:53
const (
	KJSTypeUndefined JSType = iota
	KJSTypeNull      JSType = 1
	KJSTypeBoolean   JSType = 2
	KJSTypeNumber    JSType = 3
	KJSTypeString    JSType = 4
	KJSTypeObject    JSType = 5
)

// JSTypedArrayType as declared in JavaScriptCore/JSValueRef.h:83
type JSTypedArrayType int32

// JSTypedArrayType enumeration from JavaScriptCore/JSValueRef.h:83
const (
	KJSTypedArrayTypeInt8Array         JSTypedArrayType = iota
	KJSTypedArrayTypeInt16Array        JSTypedArrayType = 1
	KJSTypedArrayTypeInt32Array        JSTypedArrayType = 2
	KJSTypedArrayTypeUint8Array        JSTypedArrayType = 3
	KJSTypedArrayTypeUint8ClampedArray JSTypedArrayType = 4
	KJSTypedArrayTypeUint16Array       JSTypedArrayType = 5
	KJSTypedArrayTypeUint32Array       JSTypedArrayType = 6
	KJSTypedArrayTypeFloat32Array      JSTypedArrayType = 7
	KJSTypedArrayTypeFloat64Array      JSTypedArrayType = 8
	KJSTypedArrayTypeArrayBuffer       JSTypedArrayType = 9
	KJSTypedArrayTypeNone              JSTypedArrayType = 10
)

const (
	// KJSPropertyAttributeNone as declared in JavaScriptCore/JSObjectRef.h:51
	KJSPropertyAttributeNone = iota
	// KJSPropertyAttributeReadOnly as declared in JavaScriptCore/JSObjectRef.h:52
	KJSPropertyAttributeReadOnly = 2
	// KJSPropertyAttributeDontEnum as declared in JavaScriptCore/JSObjectRef.h:53
	KJSPropertyAttributeDontEnum = 4
	// KJSPropertyAttributeDontDelete as declared in JavaScriptCore/JSObjectRef.h:54
	KJSPropertyAttributeDontDelete = 8
)

const (
	// KJSClassAttributeNone as declared in JavaScriptCore/JSObjectRef.h:69
	KJSClassAttributeNone = iota
	// KJSClassAttributeNoAutomaticPrototype as declared in JavaScriptCore/JSObjectRef.h:70
	KJSClassAttributeNoAutomaticPrototype = 2
)
