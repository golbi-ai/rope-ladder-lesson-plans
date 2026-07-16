<!-- rope-ladder:begin document c735ffc6fd46f09284b4950d221226070aa4c0156f62ddbd7d259e65e9db3965 -->
# Implementation-language concepts

Topics are ordered by prerequisite depth. Each repository claim links to source evidence; general facts require external evidence.

## C fixed-width integers and bitwise packing

The code uses uint8_t through uint64_t to state intended storage widths in data and bitstream paths.

Fixed-width integer values hold packed ARGB samples, codec fields, chunk tags, and arithmetic intermediates; shifts, masks, and bitwise operators extract or assemble channel values.

### How it works

#### lossless.c channel extraction

ARGB words are shifted and masked to form component values used by lossless transforms.

#### format_constants.h MKFOURCC

A four-character container tag is assembled with shifted byte values.

### Anchored code example

```
const uint32_t argb = src[i];
const int8_t green = (int8_t)(argb >> 8);
const uint32_t red = argb >> 16;
```

Source: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LTransformColor

Code examples: WebPData, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/dsp/lossless.c` — VP8LTransformColor
- Code: `thirdparty/libwebp/src/webp/format_constants.h` — MKFOURCC
- Code: `thirdparty/libwebp/src/webp/types.h` — uint8_t, uint16_t, uint32_t, uint64_t definitions
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C preprocessor feature and platform configuration

The inspected code also uses feature macros to compile TinyEXR's header implementation and platform branches in Volk.

Configuration headers use macros and conditional inclusion to enable ThorVG capabilities, select platform-dependent threading, and choose system or bundled SPIR-V headers.

### How it works

#### ThorVG feature switches

config.h enables software raster support, SVG loading, PNG loading, and conditionally thread support.

#### Header-source selection

spirv_reflect.h chooses a system SPIR-V header when SPIRV_REFLECT_USE_SYSTEM_SPIRV_H is defined.

#### Single-header implementation switch

tinyexr.cc defines TINYEXR_IMPLEMENTATION before including tinyexr.h.

### Anchored code example

```
#define THORVG_SW_RASTER_SUPPORT
#define THORVG_SVG_LOADER_SUPPORT
#define THORVG_PNG_LOADER_SUPPORT
#ifndef WEB_ENABLED
#define THORVG_THREAD_SUPPORT
#endif
```

Source: `thirdparty/thorvg/inc/config.h` — THORVG_* feature macros

Evidence:
- Code: `thirdparty/thorvg/inc/config.h` — THORVG_SW_RASTER_SUPPORT, THORVG_SVG_LOADER_SUPPORT, THORVG_PNG_LOADER_SUPPORT, THORVG_THREAD_SUPPORT
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — SPIRV_REFLECT_USE_SYSTEM_SPIRV_H
- Code: `thirdparty/tinyexr/tinyexr.cc` — TINYEXR_IMPLEMENTATION
- Code: `thirdparty/volk/volk.c` — _WIN32 and __APPLE__ preprocessor branches
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## C preprocessor portability

This is the portability layer that lets scalar and architecture-specific sources coexist.

Conditional compilation and macros select target-dependent code, inline helpers, endianness behavior, and instruction-family implementations at compile time.

### How it works

#### endian_inl_utils.h HToLE32

WORDS_BIGENDIAN selects byte-swapping macros or identity macros for host-to-little-endian conversion.

#### mips_macro.h ADD_SUB_HALVES

A macro packages repeated target assembly text while preserving operand substitution.

### Anchored code example

```
#if defined(WORDS_BIGENDIAN)
#define HToLE32 BSwap32
#define HToLE16 BSwap16
#else
#define HToLE32(x) (x)
#define HToLE16(x) (x)
#endif
```

Source: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — HToLE32 and HToLE16

Evidence:
- Code: `thirdparty/libwebp/src/utils/endian_inl_utils.h` — WORDS_BIGENDIAN, HToLE32, HToLE16
- Code: `thirdparty/libwebp/src/dsp/mips_macro.h` — ADD_SUB_HALVES
- Code: `thirdparty/libwebp/src/webp/types.h` — WEBP_INLINE conditional definition
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ standard-library module import

This is a conditional module-aware inclusion path in the generated C++ headers.

The Vulkan-Hpp façade contains an import std declaration and other headers test VULKAN_HPP_CXX_MODULE before falling back to textual includes.

### How it works

#### vulkan.hpp import std

The façade's inventory includes a standard-library module import declaration.

#### VULKAN_HPP_CXX_MODULE branch

The extension-inspection header avoids its normal includes when the Vulkan-Hpp C++ module macro is defined.

### Anchored code example

```
import std;
```

Source: `thirdparty/vulkan/include/vulkan/vulkan.hpp` — import std

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan.hpp` — import std
- Code: `thirdparty/vulkan/include/vulkan/vulkan_extension_inspection.hpp` — VULKAN_HPP_CXX_MODULE
- External (primary, verified): [C++ Working Draft: Import declaration](https://eel.is/c++draft/module.import), accessed 2026-07-16

## C++ module interface partition

The file is a C++ module interface artifact for the Vulkan video API namespace.

`vulkan_video.cppm` declares the exported `vulkan_hpp:video` module partition and an exported video namespace.

### How it works

#### Partition declaration

`export module vulkan_hpp:video;` names an exported partition of the `vulkan_hpp` module.

#### Namespace export

The file exports the namespace formed by the Vulkan-Hpp namespace macros.

### Anchored code example

```
export module vulkan_hpp:video;
export namespace VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm` — export module vulkan_hpp:video

Code examples: VideoSessionCreateInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm` — export module vulkan_hpp:video
- Code: `thirdparty/vulkan/include/vulkan/vulkan_video.cppm` — export namespace VULKAN_HPP_NAMESPACE::VULKAN_HPP_VIDEO_NAMESPACE
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C# AssemblyLoadContext plugin loading

PluginLoadContext selects shared assemblies from the main context and resolves other dependencies through AssemblyDependencyResolver.

The editor plugin loader creates AssemblyLoadContext instances to resolve dependencies and support unloading.

### How it works

#### Context construction

PluginLoadContext receives the plugin path, shared-assembly collection, main load context, and collectible flag, then creates AssemblyDependencyResolver.

#### Plugin entry loading

GodotPlugins.Main creates a PluginLoadContext, keeps a WeakReference to it, loads the requested assembly, and locates the GodotSharpEditor entry point.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/PluginLoadContext.cs` — PluginLoadContext
- Code: `modules/mono/glue/GodotSharp/GodotPlugins/Main.cs` — LoadPlugin
- External (official, unverified (source anchor not found)): [About AssemblyLoadContext - .NET](https://learn.microsoft.com/en-us/dotnet/core/dependency-loading/understanding-assemblyloadcontext), accessed 2026-07-15

## C# attributes and reflection

Godot declares attributes for scripts, exported members, signals, RPCs, tools, and serialization, then bridge code inspects attributes at runtime.

Attributes declare engine-facing metadata and reflection reads it to connect types, members, and script information.

### How it works

#### Attribute declarations

ExportAttribute constrains use to fields and properties and carries editor hint data; RpcAttribute supplies networking mode, call-local, transfer mode, and channel metadata.

#### Runtime discovery

ScriptManagerBridge reads ScriptPathAttribute, AssemblyHasScriptsAttribute, IconAttribute, RpcAttribute, and generated static metadata through reflection.

#### Analyzer inspection

GlobalClassAnalyzer uses Roslyn symbols and attributes to enforce global-class constraints during compilation.

Code examples: ScriptPathAttribute, AssemblyHasScriptsAttribute

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ExportAttribute.cs` — ExportAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/RpcAttribute.cs` — RpcAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs` — ScriptManagerBridge
- External (official, unverified (source anchor not found)): [Attributes and reflection - C#](https://learn.microsoft.com/en-us/dotnet/csharp/advanced-topics/reflection-and-attributes/), accessed 2026-07-15

## GLSL compute workgroups and synchronization

The supplied shaders use compute passes for velocity, luminance, resolve, roughness, hierarchy, skinning, and postprocessing work.

Compute shaders use declared resource interfaces, workgroup identifiers, bounds checks, shared arrays, and memory barriers for image-processing operations.

### How it works

#### Per-pixel velocity dispatch

motion_vectors_store.glsl rejects out-of-bounds invocations, derives a pixel coordinate from gl_GlobalInvocationID, and writes velocity to an image.

#### Workgroup reduction

luminance_reduce.glsl allocates shared temporary luminance storage and calls groupMemoryBarrier before reduction work.

### Anchored code example

```
if (any(greaterThanEqual(vec2(gl_GlobalInvocationID.xy), params.resolution))) {
	return;
}
```

Source: `servers/rendering/renderer_rd/shaders/effects/motion_vectors_store.glsl` — main

Prerequisites: glsl-resource-interfaces

Evidence:
- Code: `servers/rendering/renderer_rd/shaders/effects/motion_vectors_store.glsl` — main
- Code: `servers/rendering/renderer_rd/shaders/effects/luminance_reduce.glsl` — main
- Code: `servers/rendering/renderer_rd/shaders/effects/resolve.glsl` — main
- External (official, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.html), accessed 2026-07-15

## Iteration protocols

Typed and `Variant` paths are exercised for built-in and user-defined iteration.

`for` loops cover ranges, collections, and custom objects that provide iterator hook functions.

### How it works

#### Custom iterator hooks

Custom iterator classes declare `_iter_init`, `_iter_next`, and `_iter_get` methods.

#### Typed loop sources

Fixtures iterate over typed arrays, dictionaries, ranges, strings, and object iterators.

### Anchored code example

```
func _iter_get(_count) -> StringName:
```

Source: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_types.gd` — Iterator._iter_get()

Prerequisites: function-contracts

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_object.gd` — ExponentialIterator._iter_init(), _iter_next(), and _iter_get()
- Code: `modules/gdscript/tests/scripts/runtime/features/for_loop_iterator_types.gd` — Iterator._iter_get()
- Code: `modules/gdscript/tests/scripts/parser/features/for_range.gd` — test()
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

## Pattern matching and guards

Parser and runtime fixtures separately cover structural matching syntax and guarded execution.

`match` supports literals, arrays, dictionaries, bindings, multiple patterns, wildcards, and guarded branches.

### How it works

#### Structural patterns

Parser fixtures exercise array and dictionary patterns, including nested patterns and rest matching.

#### Guarded bindings

Runtime fixtures test `when` guards after bindings and verify guard-side-effect behavior.

### Anchored code example

```
var a_bind when b == 0:
```

Source: `modules/gdscript/tests/scripts/runtime/features/match_with_pattern_guards.gd` — guarded match branch

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/parser/features/match_dictionary.gd` — test()
- Code: `modules/gdscript/tests/scripts/runtime/features/match_with_pattern_guards.gd` — test()
- Code: `modules/gdscript/tests/scripts/parser/errors/match_multiple_variable_binds_in_branch.out` — expected parser error
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

## Python SCons configuration functions

Module configuration is executable build logic rather than declarative project metadata alone.

Python build files define can_build and configure functions and call SCons environment methods to select module sources and generated artifacts.

### How it works

#### Conditional module availability

The 2D physics module's can_build function reads disable_physics_2d from the build environment.

#### Source-file registration

SCsub scripts call add_source_files on the SCons environment, and the lightmapper script requests generated GLSL headers.

### Anchored code example

```
def can_build(env, platform):
    return not env["disable_physics_2d"]
```

Source: `modules/godot_physics_2d/config.py` — can_build

Evidence:
- Code: `modules/godot_physics_2d/config.py` — can_build
- Code: `modules/lightmapper_rd/SCsub` — GLSL_HEADER and add_source_files
- External (official, verified): [Python Language Reference: Function definitions](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

## C ABI structures and manual lifetime

The header is usable from C and wraps its declarations in extern "C" when compiled as C++.

The SPIR-V reflection interface exchanges C structs and pointers through create, enumerate, query, change, and explicit destroy functions.

### How it works

#### Caller-provided module record

spvReflectCreateShaderModule receives SPIR-V bytes and a caller-provided SpvReflectShaderModule pointer.

#### Explicit retirement

spvReflectDestroyShaderModule accepts a module pointer and is the paired lifetime operation.

#### Pointer-array enumeration

Enumeration functions accept a count pointer and an output array of pointers owned by the reflection module.

### Anchored code example

```
SpvReflectResult spvReflectCreateShaderModule(
  size_t                   size,
  const void*              p_code,
  SpvReflectShaderModule*  p_module
);
```

Source: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule

Code examples: SpvReflectShaderModule, SpvReflectDescriptorSet, SpvReflectDescriptorBinding

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule, spvReflectDestroyShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectEnumerateDescriptorBindings
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — extern "C" linkage guard
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## Python build hooks

Top-level and module-local scripts define functions such as can_build and configure, while SCsub files add sources and web JavaScript libraries conditionally.

Python configuration scripts define module build hooks and configure build environments.

### How it works

#### Module availability hook

The WebP module's can_build hook returns True.

#### Module configuration hook

The WebP module defines a configure hook even though its body is pass.

#### Platform-specific source configuration

WebRTC and WebXR SCsub scripts add JavaScript libraries when the platform is web.

### Anchored code example

```
def can_build(env, platform):
    return True


def configure(env):
    pass

```

Source: `modules/webp/config.py` — can_build and configure

Evidence:
- Code: `modules/webp/config.py` — can_build
- Code: `modules/webrtc/SCsub` — web platform JavaScript library configuration
- Code: `modules/webxr/SCsub` — web platform JavaScript library configuration
- External (official, verified): [The Python Language Reference: Compound statements](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

## C ABI-versioned initialization

This protects the public structure contract against caller and library version mismatch.

Public structures are initialized through inline wrappers that pass a library ABI version to internal initialization functions before callers use the fields.

### How it works

#### WebPConfigInit

The inline wrapper supplies WEBP_ENCODER_ABI_VERSION to WebPConfigInitInternal.

#### WebPPictureInit

Picture initialization similarly delegates to a version-checked internal entry point.

### Anchored code example

```
static WEBP_INLINE int WebPConfigInit(WebPConfig* config) {
  return WebPConfigInitInternal(config, WEBP_PRESET_DEFAULT, 75.f,
                                WEBP_ENCODER_ABI_VERSION);
}
```

Source: `thirdparty/libwebp/src/webp/encode.h` — WebPConfigInit

Prerequisites: c-aggregate-contracts

Code examples: WebPConfig, WebPPicture

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — WEBP_ENCODER_ABI_VERSION, WebPConfigInit, WebPPictureInit
- Code: `thirdparty/libwebp/src/enc/config_enc.c` — WebPConfigInitInternal
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C macro-based binary decoding

The MiniUPnP decoder offers pointer, callback-reader, and bounds-checked variants of the same 7-bit encoding loop.

C macros implement variable-length decoding by shifting accumulated values and testing continuation bits.

### How it works

#### DECODELENGTH

The macro advances a byte pointer while the high continuation bit is set.

#### DECODELENGTH_READ

The callback form obtains each input byte through a supplied reader macro.

#### DECODELENGTH_CHECKLIMIT

The bounded form checks the input limit before consuming data.

### Anchored code example

```
n = (n << 7) | (c & 0x07f);
```

Source: `thirdparty/miniupnpc/src/codelength.h` — DECODELENGTH_READ

Prerequisites: c-conditional-compilation

Evidence:
- Code: `thirdparty/miniupnpc/src/codelength.h` — DECODELENGTH
- Code: `thirdparty/miniupnpc/src/codelength.h` — DECODELENGTH_CHECKLIMIT
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C parser state machines

The XML parser state stores the start, end, and current input positions for the parsing routines.

C structs retain XML cursor state and parser result heads while miniUPnPc parsing functions build protocol records.

### How it works

#### xmlparser

The parser record tracks the source boundaries and current character pointer.

#### NameValueParserData

Name/value parsing maintains a linked result head.

#### PortMappingParserData

Port-list parsing maintains a linked mapping head.

### Anchored code example

```
struct xmlparser {
const char *xmlstart;
const char *xmlend;
const char *xml;	/* pointer to current character */
```

Source: `thirdparty/miniupnpc/src/minixml.h` — struct xmlparser

Prerequisites: c-struct-linked-state

Code examples: NameValueParserData, PortMappingParserData

Evidence:
- Code: `thirdparty/miniupnpc/src/minixml.h` — struct xmlparser
- Code: `thirdparty/miniupnpc/include/miniupnpc/upnpreplyparse.h` — NameValueParserData
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C preprocessor composition

FreeType uses this source-level composition pattern for both its base and autofit module entry translation units.

C preprocessing defines FT_MAKE_OPTION_SINGLE_OBJECT and then includes component implementation files to form a single-object library unit.

### How it works

#### Autofit aggregation

autofit.c defines FT_MAKE_OPTION_SINGLE_OBJECT before including the autofit implementation files.

#### Base aggregation

ftbase.c uses the same macro before including the base implementation files.

### Anchored code example

```
#define FT_MAKE_OPTION_SINGLE_OBJECT

#include "ft-hb.c"
#include "ft-hb-ft.c"
#include "afadjust.c"
```

Source: `thirdparty/freetype/src/autofit/autofit.c` — FT_MAKE_OPTION_SINGLE_OBJECT

Evidence:
- Code: `thirdparty/freetype/src/autofit/autofit.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- Code: `thirdparty/freetype/src/base/ftbase.c` — FT_MAKE_OPTION_SINGLE_OBJECT
- External (primary, unverified (source anchor not found)): [N1570 — Committee Draft, Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C structures and pointer handles

FreeType records commonly embed a base record and keep state behind pointer typedefs.

C structures and pointers package mutable subsystem state and expose opaque handle types.

### How it works

#### Bzip2 stream state

FT_BZip2FileRec stores source and embedding streams, memory, bzlib state, buffers, cursors, and reset status in one record.

#### PFR face extension

PFR_FaceRec embeds FT_FaceRec as root and adds parsed PFR header, logical-font, and physical-font data.

#### SVG renderer extension

SVG_RendererRec embeds FT_RendererRec as root and adds hooks plus hook-owned state.

### Anchored code example

```
typedef struct  FT_BZip2FileRec_
{
  FT_Stream  source;         /* parent/source stream        */
  FT_Stream  stream;         /* embedding stream            */
  FT_Memory  memory;         /* memory allocator            */
  bz_stream  bzstream;       /* bzlib input stream          */

  FT_Byte    input[FT_BZIP2_BUFFER_SIZE];  /* input read buffer  */

  FT_Byte    buffer[FT_BZIP2_BUFFER_SIZE]; /* output buffer          */
  FT_ULong   pos;                          /* position in output     */
  FT_Byte*   cursor;
  FT_Byte*   limit;
  FT_Bool    reset;                        /* reset before next read */

} FT_BZip2FileRec, *FT_BZip2File;
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — FT_BZip2FileRec_

Code examples: FT_BZip2FileRec, PFR_FaceRec, SVG_RendererRec

Evidence:
- Code: `thirdparty/freetype/src/pfr/pfrobjs.h` — PFR_FaceRec_
- Code: `thirdparty/freetype/src/svg/svgtypes.h` — SVG_RendererRec_
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ native class hierarchy

The engine’s native implementation expresses its core object, scene, resource, extension, and scripting types as C++ classes, then exposes selected APIs through ClassDB.

Native components use C++ classes, inheritance, and virtual functions to implement engine types such as Resource, Node, and GDScript.

### How it works

#### Resource inheritance

Resource derives from RefCounted, and PackedScene derives from Resource, establishing the resource specialization chain.

#### Virtual subsystem APIs

Mesh declares virtual surface and material operations so concrete mesh types can supply their implementations.

#### Native registration

ClassDB registration stores creation functions and exposure metadata for registered concrete classes.

### Anchored code example

```
class PackedScene : public Resource {
	GDCLASS(PackedScene, Resource);
```

Source: `scene/resources/packed_scene.h` — PackedScene

Code examples: Resource, Node, PackedScene, GDScript

Evidence:
- Code: `scene/resources/packed_scene.h` — PackedScene
- Code: `core/object/class_db.h` — ClassDB::register_class
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

## C dynamic-library function wrappers

The generated wrapper headers declare replacement function-pointer symbols and identify the wrapped shared-library sonames.

Generated C function-pointer wrappers dynamically route DBus, Fontconfig, Speech Dispatcher, and Wayland/libdecor calls for Linux/BSD desktop integration.

### How it works

#### dbus-so_wrap.c

The DBus wrapper declares function pointers matching DBus API signatures.

#### fontconfig-so_wrap.h

The Fontconfig wrapper header identifies itself as generated and redirects API names through wrapper symbols.

#### wayland-client-core-so_wrap.c

The Wayland wrapper declares function pointers for core Wayland client calls such as display connection and event-queue creation.

### Anchored code example

```
const char* (*dbus_address_entry_get_value_dylibloader_wrapper_dbus)( DBusAddressEntry*,const char*);
```

Source: `platform/linuxbsd/dbus-so_wrap.c` — dbus_address_entry_get_value_dylibloader_wrapper_dbus

Evidence:
- Code: `platform/linuxbsd/dbus-so_wrap.c` — dbus_address_entry_get_value_dylibloader_wrapper_dbus
- Code: `platform/linuxbsd/fontconfig-so_wrap.h` — DYLIBLOAD_WRAPPER_FONTCONFIG
- Code: `platform/linuxbsd/wayland/dynwrappers/wayland-client-core-so_wrap.c` — wl_display_connect_dylibloader_wrapper_wayland_client
- External (official, unverified (source anchor not found)): [ISO/IEC 9899:2024 - Programming languages — C](https://www.iso.org/standard/82075.html), accessed 2026-07-15

## C pointers, arrays, and strides

Most image operations accept borrowed pointer ranges and compute row addresses instead of owning per-pixel objects.

C aggregates are reached through pointers and stepped across image memory with width- and stride-based pointer arithmetic for rows and planes.

### How it works

#### VP8IteratorImport

The iterator derives luma and chroma source addresses from plane bases, macroblock coordinates, and distinct strides.

#### WebPDecodeYUV

The decoding API returns luma and chroma pointers plus separate luma and UV strides.

### Anchored code example

```
const uint8_t* const ysrc = pic->y + (y * pic->y_stride  + x) * 16;
const uint8_t* const usrc = pic->u + (y * pic->uv_stride + x) * 8;
const uint8_t* const vsrc = pic->v + (y * pic->uv_stride + x) * 8;
```

Source: `thirdparty/libwebp/src/enc/iterator_enc.c` — VP8IteratorImport

Prerequisites: c-aggregate-contracts

Code examples: WebPPicture, WebPRescaler, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/enc/iterator_enc.c` — VP8IteratorImport
- Code: `thirdparty/libwebp/src/webp/decode.h` — WebPDecodeYUV
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C-compatible function-pointer interfaces

This is a data-driven dispatch mechanism used alongside C++ class polymorphism.

C-compatible function-pointer interfaces let converter implementations select operations through a shared implementation record.

### How it works

#### UConverterSharedData::impl

Shared converter data points to an implementation record documented as a vtable-style group of function pointers.

#### UDataSwapper

The swapper structure contains function pointers for binary-data transformation operations.

### Anchored code example

```
const UConverterImpl *impl;     /* vtable-style struct of mostly function pointers */
```

Source: `thirdparty/icu4c/common/ucnv_bld.h` — UConverterSharedData::impl

Prerequisites: cpp-pointers-and-const

Code examples: UConverter, UConverterSharedData

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h` — UConverterSharedData::impl
- Code: `thirdparty/icu4c/common/udataswp.h` — UDataSwapper
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## C++ object-representation casts

These operations are concentrated in code that consumes serialized binary layouts.

Pointer reinterpretation reads packed data through typed views, particularly when ICU validates or loads binary resource and Unicode data.

### How it works

#### SpoofImpl data loading

Spoof implementation code derives a UDataInfo pointer by viewing bytes at an offset.

#### UCPTrie opening

UCPTrie code views serialized trie storage through typed pointers.

### Anchored code example

```
const UDataInfo *pInfo = (const UDataInfo *)((const char *)inData+4);
```

Source: `thirdparty/icu4c/i18n/uspoof_impl.cpp` — SpoofImpl data loader

Prerequisites: cpp-pointers-and-const

Code examples: UResourceBundle, UCPTrie

Evidence:
- Code: `thirdparty/icu4c/i18n/uspoof_impl.cpp` — SpoofImpl data loader
- Code: `thirdparty/icu4c/common/ucptrie.cpp` — ucptrie_openFromBinary
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## Python build-action functions

Python is used for editor build support, not the editor’s primary runtime implementation.

Python defines the icon-build action that receives target, source, and environment arguments, while SCsub scripts invoke Python-based build-environment methods.

### How it works

#### Build action definition

make_editor_icons_action is a Python function with target, source, and env parameters.

#### Build-script calls

SCsub scripts call env.add_source_files to contribute C++ files to env.editor_sources.

### Anchored code example

```
def make_editor_icons_action(target, source, env):
```

Source: `editor/icons/editor_icons_builders.py` — make_editor_icons_action

Evidence:
- Code: `editor/icons/editor_icons_builders.py` — make_editor_icons_action
- Code: `editor/import/SCsub` — env.add_source_files
- External (official, verified): [Python documentation — Defining Functions](https://docs.python.org/3/tutorial/controlflow.html#defining-functions), accessed 2026-07-16

## C aggregate state and callback modules

The implementation uses public context structures plus private extension structures instead of language-level objects.

C aggregate types and typed function pointers define stateful processing modules, letting JPEG upsamplers retain shared buffers and dispatch a per-component routine.

### How it works

#### jdsample.h: upsample1_ptr

The typedef fixes the signature for a per-component upsampling routine, including the decompression context, component metadata, and input/output sample arrays.

#### jdsample.h: my_upsampler

The private upsampler aggregates color-row buffers, an array of routine pointers, row counters, and cached expansion factors.

#### jpeglib.h: jpeg_decompress_struct

The public decompression context stores pointers to the selected decompression modules, including inverse DCT, upsampling, color conversion, and quantization.

### Anchored code example

```
typedef void (*upsample1_ptr) (j_decompress_ptr cinfo,
                               jpeg_component_info *compptr,
                               _JSAMPARRAY input_data,
                               _JSAMPARRAY *output_data_ptr);
```

Source: `thirdparty/libjpeg-turbo/src/jdsample.h` — upsample1_ptr

Code examples: JPEG Decompression Context, JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdsample.h` — upsample1_ptr
- Code: `thirdparty/libjpeg-turbo/src/jdsample.h` — my_upsampler
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming Languages: C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C resource lifecycles and ownership

Resource retirement is explicit; structure storage itself is generally caller-owned.

Pointer-bearing structures use explicit init, allocate, clear, free, copy, and view functions so the caller can distinguish owned image buffers from borrowed views.

### How it works

#### WebPPictureAlloc and WebPPictureFree

The picture API allocates or releases pixel memory while retaining the caller's structure object.

#### WebPMemoryWriterInit and WebPMemoryWriterClear

The memory writer initializes and later releases its owned output allocation.

Prerequisites: c-aggregate-contracts, c-pointers-arrays-and-strides

Code examples: WebPPicture, WebPMemoryWriter, WebPDecBuffer

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — WebPPictureAlloc, WebPPictureFree, WebPPictureView, WebPMemoryWriterClear
- Code: `thirdparty/libwebp/src/enc/picture_enc.c` — WebPPictureAlloc
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ Basis transcoding behind a C-compatible API

This is the material C++ portion of the supplied KTX implementation; most neighboring public and internal KTX APIs are C declarations.

The KTX Basis transcode implementation uses C++ references and standard-library state vectors, while its headers retain a C-compatible interface for the enclosing C API.

### How it works

#### basis_transcode.cpp: ktxTexture2_transcodeEtc1s

The ETC1S path binds each image description as a C++ const reference before processing it.

#### basis_transcode.cpp: ktxTexture2_transcodeUastc

The UASTC path constructs a Basis Universal transcoder and maintains transcoder state in a std::vector.

#### basis_sgd.h: __cplusplus linkage guard

The shared global-data header surrounds its C declarations with an extern "C" guard when included from C++.

### Anchored code example

```
const ktxBasisLzEtc1sImageDesc& imageDesc = imageDescs[image];
```

Source: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeEtc1s

Prerequisites: c-aggregate-callback-modules

Code examples: KTX2 Texture, KTX2 Private Texture State

Evidence:
- Code: `thirdparty/libktx/lib/basis_transcode.cpp` — ktxTexture2_transcodeEtc1s
- Code: `thirdparty/libktx/lib/basis_sgd.h` — __cplusplus
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4861](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4861.pdf), accessed 2026-07-15

## C++ runtime class registration

Godot's native API surface is assembled through ClassDB rather than by a separate hand-written runtime schema.

C++ templates and classes implement the runtime registry that records native class identity, inheritance, construction, methods, properties, signals, and constants.

### How it works

#### Class registration

ClassDB::register_class initializes the type metadata, assigns a creator function, marks the class exposed, and records the current API category.

#### Method binding

ClassDB::bind_method builds a MethodBind from a C++ member function and stores it under an exposed method name.

#### Property and signal metadata

ClassDB::add_property and ClassDB::add_signal populate the reflection data later consumed by the engine API.

### Anchored code example

```
static bool class_exists(const StringName &p_class);
static bool is_parent_class(const StringName &p_class, const StringName &p_inherits);
static bool can_instantiate(const StringName &p_class);
```

Source: `core/object/class_db.h` — ClassDB

Code examples: Resource, RID

Evidence:
- Code: `core/object/class_db.h` — ClassDB
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ inheritance and engine reference handles

The supplied C++ code combines class hierarchies with engine-specific ownership wrappers in tests and mocks.

C++ test infrastructure uses public inheritance for test doubles and stores engine objects in Ref<T> handles created with memnew.

### How it works

#### Display-server test double

DisplayServerMock publicly derives from DisplayServerHeadless, making the mock usable where the headless display-server interface is expected.

#### Signal observation object

SignalWatcher publicly derives from Object, placing the watcher in the engine object hierarchy.

#### Animation resource handle

The animation tests bind a memnew(Animation) result to Ref<Animation> before adding tracks.

### Anchored code example

```
class DisplayServerMock : public DisplayServerHeadless {
```

Source: `tests/display_server_mock.h` — DisplayServerMock declaration

Code examples: Animation, AnimationTrack

Evidence:
- Code: `tests/display_server_mock.h` — DisplayServerMock
- Code: `tests/signal_watcher.h` — SignalWatcher
- Code: `tests/scene/test_animation.cpp` — Ref<Animation> allocation
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2024/n4981.pdf), accessed 2026-07-15

## C++ enumerated export state

Named integral categories keep export configuration and reporting states explicit at API boundaries.

C++ enumerations classify export messages, export filters, and file or script export modes.

### How it works

#### Message severity

EditorExportPlatform::ExportMessageType classifies information, warnings, and errors emitted during export.

#### Preset resource policy

EditorExportPreset::ExportFilter and FileExportMode encode file-selection and per-file customization policies.

#### Script representation

EditorExportPreset::ScriptExportMode encodes text and binary script export choices.

### Anchored code example

```
enum ExportMessageType {
```

Source: `editor/export/editor_export_platform.h` — EditorExportPlatform::ExportMessageType

Code examples: EditorExportPreset, EditorExportPlatform::ExportMessage

Evidence:
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::ExportMessageType
- Code: `editor/export/editor_export_preset.h` — EditorExportPreset::ExportFilter
- Code: `editor/export/editor_export_preset.h` — EditorExportPreset::ScriptExportMode
- External (primary, unverified (source anchor not found)): [Draft C++ Standard — Enumerations](https://eel.is/c++draft/dcl.enum), accessed 2026-07-15

## C++ binary data codecs

The accessor implementation separates byte transport from conversion into numeric and vector values.

glTF decoding reads typed C++ containers of raw bytes through buffer-view offsets and strides, while encoding writes byte arrays into newly allocated buffer views.

### How it works

#### Buffer-view byte extraction

GLTFBufferView obtains bytes from the indexed buffer and calculates an end offset from byte_offset and byte_length.

#### Accessor byte decoding

GLTFAccessor loads a PackedByteArray from a buffer view before interpreting component-sized values.

#### Encoded buffer-view creation

Accessor encoding sends generated bytes to GLTFBufferView::write_new_buffer_view_into_state.

### Anchored code example

```
const PackedByteArray raw_bytes = raw_buffer_view->load_buffer_view_data(p_gltf_state);
```

Source: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers

Prerequisites: cpp-resource-containers

Code examples: GLTFBufferView, GLTFAccessor

Evidence:
- Code: `modules/gltf/structures/gltf_buffer_view.cpp` — GLTFBufferView::load_buffer_view_data
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor byte decode and buffer-view write paths
- External (primary, verified): [C++ Working Draft: Fundamental types](https://eel.is/c++draft/basic.types), accessed 2026-07-16

## C++ class state and composition

The GamepadMotion helper is implemented in a C++ header containing class and struct definitions plus their behavior.

C++ classes group motion settings, calibration, vector, quaternion, and motion-update behavior behind named types.

### How it works

#### Motion domain types

GamepadMotion.hpp declares GyroCalibration, Quat, Vec, SensorMinMaxWindow, AutoCalibration, Motion, GamepadMotionSettings, and GamepadMotion.

#### Stateful motion API

GamepadMotion is the top-level stateful type, while Motion and AutoCalibration separate orientation update and calibration responsibilities.

### Anchored code example

```
class GamepadMotion
```

Source: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — GamepadMotion

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — GamepadMotionSettings
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Motion
- External (primary, unverified (source anchor not found)): [N4950: Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ classes and inheritance

Class declarations organize most ICU and Jolt stateful services; HarfBuzz more often uses structs.

C++ classes define reusable state and behavior, and the implementation derives UnicodeString from Replaceable and many ICU services from UObject or UMemory.

### How it works

#### UnicodeString

UnicodeString is declared as a derived class, making Replaceable operations part of its type interface.

#### LanguageBreakEngine

LanguageBreakEngine is a UObject-derived base class used by multiple break-engine implementations.

### Anchored code example

```
class U_COMMON_API UnicodeString : public Replaceable
```

Source: `thirdparty/icu4c/common/unicode/unistr.h` — UnicodeString

Code examples: Locale, AABBTreeBuilder

Evidence:
- Code: `thirdparty/icu4c/common/unicode/unistr.h` — UnicodeString
- Code: `thirdparty/icu4c/common/brkeng.h` — LanguageBreakEngine
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## C++ classes and inheritance

Class declarations establish the editor subsystem types and their engine-facing base behavior.

The editor source declares engine services as C++ classes derived from Object, Resource, Control, and other engine base classes.

### How it works

#### Persisted settings resource

EditorSettings derives from Resource, placing editor settings in the engine's resource hierarchy.

#### Interactive viewport control

Node3DEditorViewport derives from Control, making the 3D editor viewport a GUI control type.

#### Editor-facing object interface

EditorVCSInterface derives from Object and groups version-control data structures with its interface.

### Anchored code example

```
class EditorSettings : public Resource {
```

Source: `editor/settings/editor_settings.h` — EditorSettings

Code examples: EditorSettings, VCS Diff File

Evidence:
- Code: `editor/settings/editor_settings.h` — EditorSettings : public Resource
- Code: `editor/scene/3d/node_3d_editor_viewport.h` — Node3DEditorViewport : public Control
- Code: `editor/version_control/editor_vcs_interface.h` — EditorVCSInterface : public Object
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ exception containment at ABI boundaries

The macros are attached to exported loader functions so exceptions do not escape the loader ABI.

OpenXR loader macros use C++ exceptions internally but convert `std::bad_alloc` and other `std::exception` failures into OpenXR result codes.

### How it works

#### XRLOADER_ABI_TRY

The macro expands to a C++ `try` block unless exception handling is disabled.

#### XRLOADER_ABI_CATCH_BAD_ALLOC_OOM

Allocation failure is logged and mapped to `XR_ERROR_OUT_OF_MEMORY`.

#### XRLOADER_ABI_CATCH_FALLBACK

Other standard exceptions are logged and mapped to `XR_ERROR_RUNTIME_FAILURE`.

### Anchored code example

```
#define XRLOADER_ABI_TRY try
```

Source: `thirdparty/openxr/src/loader/exception_handling.hpp` — XRLOADER_ABI_TRY

Code examples: LoaderInstance

Evidence:
- Code: `thirdparty/openxr/src/loader/exception_handling.hpp` — XRLOADER_ABI_CATCH_BAD_ALLOC_OOM
- Code: `thirdparty/openxr/src/loader/exception_handling.hpp` — XRLOADER_ABI_CATCH_FALLBACK
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ polymorphic backend classes

The repository uses public inheritance across audio, filesystem, networking, rendering-context, rendering-device, and editor plugin implementations.

C++ backend classes inherit engine interfaces so a rendering or platform service can be selected through a common base type.

### How it works

#### Rendering context specialization

RenderingContextDriverVulkan derives from RenderingContextDriver, identifying a concrete context implementation.

#### Rendering device specialization

RenderingDeviceDriverVulkan derives from RenderingDeviceDriver, and the Direct3D 12 and Metal drivers use the same abstraction pattern.

#### Platform service specialization

AudioDriverALSA derives from AudioDriver, while Unix file and socket implementations derive their corresponding engine interfaces.

Code examples: RenderingDeviceDriverVulkan::BufferInfo

Evidence:
- Code: `drivers/vulkan/rendering_context_driver_vulkan.h` — class RenderingContextDriverVulkan : public RenderingContextDriver
- Code: `drivers/alsa/audio_driver_alsa.h` — class AudioDriverALSA : public AudioDriver
- External (primary, unverified (source anchor not found)): [C++ Working Draft: Virtual Functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-15

## C++ resources and polymorphic engine objects

The implementation derives persisted configuration from Resource and runtime services from engine base classes such as MultiplayerAPI and NavigationServer2D or NavigationServer3D.

C++ engine modules define polymorphic Resource and server objects that own long-lived configuration and runtime state.

### How it works

#### Persisted replication resource

SceneReplicationConfig derives from Resource, declares a save type and resource extension, and maintains property lists and replication modes.

#### Runtime service inheritance

SceneMultiplayer derives from MultiplayerAPI, while the 2D and 3D servers derive from their respective NavigationServer base classes.

#### Specialized resource hierarchies

Noise, OggPacketSequence, OpenXRActionMap, and related types use Resource inheritance for engine-facing configuration and content.

### Anchored code example

```
class SceneReplicationConfig : public Resource {
	GDCLASS(SceneReplicationConfig, Resource);
	OBJ_SAVE_TYPE(SceneReplicationConfig);
```

Source: `modules/multiplayer/scene_replication_config.h` — SceneReplicationConfig

Code examples: SceneReplicationConfig, OggPacketSequence, OpenXRActionMap

Evidence:
- Code: `modules/multiplayer/scene_replication_config.h` — SceneReplicationConfig
- Code: `modules/multiplayer/scene_multiplayer.h` — SceneMultiplayer
- Code: `modules/navigation_2d/2d/godot_navigation_server_2d.h` — GodotNavigationServer2D
- Code: `modules/navigation_3d/3d/godot_navigation_server_3d.h` — GodotNavigationServer3D
- External (primary, verified): [C++ working draft — Object model](https://eel.is/c++draft/intro.object), accessed 2026-07-16

## C++ inheritance and reference-counted adapters

The Android wrapper types and display/OS classes use C++ classes and inheritance as the native adaptation mechanism.

C++ defines reference-counted Java adapter classes and platform subclasses that specialize common engine interfaces.

### How it works

#### JavaClass

`JavaClass` is declared as a `RefCounted` class, establishing a reference-counted native wrapper type.

#### JavaObject

`JavaObject` is another `RefCounted` wrapper class in the same Android API header.

#### DisplayServerAndroid

`DisplayServerAndroid` derives from `DisplayServer`, placing Android display behavior behind the engine display-server interface.

### Anchored code example

```
class JavaClass : public RefCounted {
```

Source: `platform/android/api/java_class_wrapper.h` — JavaClass

Evidence:
- Code: `platform/android/api/java_class_wrapper.h` — JavaClass
- Code: `platform/android/api/java_class_wrapper.h` — JavaObject
- Code: `platform/android/display_server_android.h` — DisplayServerAndroid
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

## C++ scope-bound locking

The repeated local MutexLock declarations make the synchronization boundary visible at write sites.

The Jolt contact listener creates scope-bound C++ lock objects before mutating shared contact, manifold, and debug-contact collections.

### How it works

#### Contact write protection

JoltContactListener3D creates MutexLock write_lock(write_mutex) before handling contact data.

#### Shared debug-contact updates

The listener uses the same lock pattern around debug-contact capacity and append operations.

### Anchored code example

```
const MutexLock write_lock(write_mutex);
```

Source: `modules/jolt_physics/spaces/jolt_contact_listener_3d.cpp` — JoltContactListener3D contact write paths

Code examples: JoltSpace3D

Evidence:
- Code: `modules/jolt_physics/spaces/jolt_contact_listener_3d.cpp` — MutexLock write_lock
- External (primary, verified): [C++ Working Draft: Mutex requirements](https://eel.is/c++draft/thread.mutex.requirements), accessed 2026-07-16

## C++ typed API records

The C++ binding exposes API data as named records instead of raw untyped parameter lists.

Generated C++ structs use typed pointers and brace defaults to represent Vulkan creation inputs and optional state records.

### How it works

#### Pipeline state composition

`GraphicsPipelineCreateInfo` has typed pointers to shader-stage and fixed-function state records.

#### Default-initialized pointers

Pointer members in the generated records are initialized with `{}`, representing an unset pointer value in the declaration.

#### Presentation arrays

`PresentInfoKHR` carries typed pointers for wait semaphores, swapchains, and image indices.

### Anchored code example

```
const PipelineShaderStageCreateInfo *        pStages             = {};
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — GraphicsPipelineCreateInfo::pStages

Code examples: GraphicsPipelineCreateInfo, PresentInfoKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct GraphicsPipelineCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — struct PresentInfoKHR
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ atomic synchronization

The implementation wraps std::atomic and uses it directly in barrier and allocator-adjacent synchronization paths.

C++ atomic counters and compare-exchange loops coordinate active barriers and conditional minimum or maximum updates across workers.

### How it works

#### Atomic wrapper

embree::atomic derives from std::atomic<T> and defines copy construction and assignment through load and store.

#### Compare-exchange update

_atomic_min and _atomic_max repeatedly load, compare, and compare_exchange_strong until the requested bound is installed or already satisfied.

#### Active barrier

BarrierActive increments an atomic counter and spins until it equals the required thread count.

### Anchored code example

```
const T b = bref.load();
    while (true) {
      T a = aref.load();
      if (a <= b) break;
      if (aref.compare_exchange_strong(a,b)) break;
    }
```

Source: `thirdparty/embree/common/sys/atomic.h` — _atomic_min

Code examples: atomic, BarrierActive, BarrierActiveAutoReset

Evidence:
- Code: `thirdparty/embree/common/sys/atomic.h` — _atomic_min
- Code: `thirdparty/embree/common/sys/atomic.h` — atomic
- Code: `thirdparty/embree/common/sys/barrier.h` — BarrierActive
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ linkage and opaque API handles

Implementation classes remain internal while API calls exchange opaque handle types and plain structures such as RTCRayHit.

C++ exposes RTCDevice, RTCScene, RTCGeometry, and RTCBuffer as opaque pointer handles through a C-linkage public boundary.

### How it works

#### Opaque handles

The public headers typedef incomplete struct pointers for device, scene, geometry, and buffer handles.

#### Linkage configuration

rtcore_config.h defines API linkage macros that choose extern C or extern C++ depending on API namespace configuration.

#### Plain query structures

rtcore_ray.h exposes public ray and hit structs independently of internal C++ RayK and HitK representations.

### Anchored code example

```
typedef struct RTCDeviceTy* RTCDevice;
typedef struct RTCSceneTy* RTCScene;
```

Source: `thirdparty/embree/include/embree4/rtcore_device.h` — RTCDevice

Code examples: RTCDevice, RTCScene, RTCGeometry, RTCBuffer

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h` — RTCDevice
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h` — RTCBuffer
- Code: `thirdparty/embree/include/embree4/rtcore_config.h` — RTC_API_EXTERN_C
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ preprocessor configuration

The configuration is compile-time: the resulting build exposes only code paths enabled by its macro definitions.

C++ preprocessing selects tasking backends, supported geometry features, ISA namespaces, platform headers, and fallback shims before compilation.

### How it works

#### ISA selection

sysinfo.h maps AVX-512, AVX2, AVX, and SSE feature macros to an isa namespace and ISA constant.

#### Feature gates

config.h defines or omits EMBREE_GEOMETRY_* and related macros, then maps them to IF_ENABLED_* expansion macros.

#### Task backend selection

taskscheduler.h includes an internal, TBB, or PPL scheduler header according to TASKING_* macros.

### Anchored code example

```
#if defined (__AVX512VL__)
#  define isa avx512
#  define ISA AVX512
#  define ISA_STR "AVX512"
#elif defined (__AVX2__)
#  define isa avx2
#  define ISA AVX2
```

Source: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros

Code examples: TaskScheduler, Geometry, vint8

Evidence:
- Code: `thirdparty/embree/common/sys/sysinfo.h` — ISA selection macros
- Code: `thirdparty/embree/kernels/config.h` — IF_ENABLED_TRIS
- Code: `thirdparty/embree/common/tasking/taskscheduler.h` — TASKING_INTERNAL
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C ABI bridging

Calling-convention macros are selected per platform and are used in both exported functions and function-pointer declarations.

OpenXR headers use macros and `extern "C"` guards so C++ callers expose C-compatible API declarations and function-pointer types.

### How it works

#### XRAPI_CALL

The platform header selects a calling convention, including `__stdcall` on Windows.

#### XRAPI_PTR

The same calling-convention selection is applied to API function-pointer types.

#### extern "C" guard

The header opens C linkage when included by C++ code.

### Anchored code example

```
#ifdef __cplusplus
extern "C" {
#endif
```

Source: `thirdparty/openxr/include/openxr/openxr_platform_defines.h` — C linkage guard

Prerequisites: c-conditional-compilation

Code examples: XrInstanceCreateInfo

Evidence:
- Code: `thirdparty/openxr/include/openxr/openxr_platform_defines.h` — XRAPI_CALL
- Code: `thirdparty/openxr/include/openxr/openxr_platform_defines.h` — XRAPI_PTR
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C ABI structures and opaque state

Vulkan descriptors expose typed fields, while zlib hides its internal state behind a pointer.

C declarations use tagged structures, pointer members, and opaque implementation pointers to express caller-owned descriptors and library-managed state across API boundaries.

### How it works

#### Caller-owned Vulkan descriptor

VkWaylandSurfaceCreateInfoKHR carries structure type, extension pointer, flags, display, and surface fields into a creation call.

#### Opaque compression state

zlib.h declares struct internal_state and stores its pointer in the public stream state rather than exposing the implementation layout.

#### Forward-declared implementation contexts

Wslay and Zstandard headers declare context structures used through API-level pointers.

### Anchored code example

```
typedef struct VkWaylandSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkWaylandSurfaceCreateFlagsKHR    flags;
    struct wl_display*                display;
    struct wl_surface*                surface;
} VkWaylandSurfaceCreateInfoKHR;
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — VkWaylandSurfaceCreateInfoKHR

Code examples: VkWaylandSurfaceCreateInfoKHR, z_stream

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — VkWaylandSurfaceCreateInfoKHR
- Code: `thirdparty/zlib/zlib.h` — struct internal_state FAR *state
- Code: `thirdparty/wslay/wslay/wslay.h` — struct wslay_frame_context and struct wslay_event_context
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C structs and pointer-linked state

SDL uses named and anonymous structs together with pointers for queues, device records, property storage, and platform handles.

C structs and pointers represent explicit runtime state, ownership links, and opaque-handle implementation data.

### How it works

#### SDL_EventEntry in SDL_events.c

Each queued event embeds an SDL_Event, points at temporary payload memory, and links to previous and next queue entries.

#### SDL_EventQ in SDL_events.c

The queue state groups its lock, active flag, atomic count, head, tail, and reusable-entry list in one anonymous C struct.

#### hid_device_info in hidapi.h

The HID enumeration API returns linked device-information records through a next pointer.

### Anchored code example

```
typedef struct SDL_EventEntry
{
    SDL_Event event;
    SDL_TemporaryMemory *memory;
    struct SDL_EventEntry *prev;
    struct SDL_EventEntry *next;
} SDL_EventEntry;
```

Source: `thirdparty/sdl/events/SDL_events.c` — SDL_EventEntry

Code examples: SDL_EventEntry, SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c` — SDL_EventEntry
- Code: `thirdparty/sdl/hidapi/hidapi/hidapi.h` — hid_device_info
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## C++ classes and inheritance

Class inheritance provides the common contracts on which the editor’s services and UI controls are built.

The editor is structured from C++ classes that derive from engine base types such as Node, Container, ScrollContainer, ResourceImporter, and RefCounted.

### How it works

#### Editor plug-in base

EditorPlugin derives from Node, making plug-ins engine nodes with editor-specific hooks.

#### Importer base

ResourceImporterScene derives from ResourceImporter, separating scene-specific importing from the generic resource-import contract.

#### Inspector UI bases

EditorProperty and EditorInspector derive from GUI container classes, so property editors compose into the editor UI hierarchy.

### Anchored code example

```
class EditorPlugin : public Node {
```

Source: `editor/plugins/editor_plugin.h` — EditorPlugin

Evidence:
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin
- Code: `editor/import/3d/resource_importer_scene.h` — ResourceImporterScene
- Code: `editor/inspector/editor_inspector.h` — EditorProperty
- External (primary, verified): [C++ Working Draft — Derived classes](https://eel.is/c++draft/class.derived), accessed 2026-07-16

## C++ flag stringification

It is a diagnostic representation helper for the generated C++ Vulkan binding.

The helper converts typed `FormatFeatureFlags` values into a string by testing individual flag bits and appending their names.

### How it works

#### Bitwise flag test

Each conditional tests whether a named `FormatFeatureFlagBits` enumerator is present in the supplied flag value.

#### Incremental textual output

The function initializes a result string and appends a name for each matching flag.

#### Library feature selection

The header checks `__cpp_lib_format` and includes either `<format>` or `<sstream>`.

### Anchored code example

```
if ( value & FormatFeatureFlagBits::eSampledImage )
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp` — to_string( FormatFeatureFlags value )

Prerequisites: cpp-typed-api-records

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp` — to_string( FormatFeatureFlags value )
- Code: `thirdparty/vulkan/include/vulkan/vulkan_to_string.hpp` — __cpp_lib_format
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ numeric value operations

The motion helper uses float-valued vectors and quaternions together with standard-library clamps and maxima.

C++ member-based value types calculate vector lengths, normalized directions, quaternion orientation, and time-scaled motion updates.

### How it works

#### Angular integration inputs

Motion update constructs an axis vector from gyroscope samples, computes angular speed, and scales the angle by delta time.

#### Gravity correction

The implementation normalizes acceleration and gravity vectors, bounds dot products, and uses correction speeds from settings.

### Anchored code example

```
const Vec axis = Vec(inGyroX, inGyroY, inGyroZ);
const Vec accel = Vec(inAccelX, inAccelY, inAccelZ);
const float angleSpeed = axis.Length() * (float)M_PI / 180.0f;
const float angle = angleSpeed * deltaTime;
```

Source: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Motion

Prerequisites: cpp-class-state-and-composition

Evidence:
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Quat
- Code: `thirdparty/gamepadmotionhelpers/GamepadMotion.hpp` — Vec
- External (primary, unverified (source anchor not found)): [N4950: Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ member-function overloading

The overload set preserves one conceptual operation while accepting multiple input representations.

Overloaded member functions provide related operations for different argument forms, such as LocaleMatcher matching a single locale, an iterator, or a list string.

### How it works

#### LocaleMatcher::getBestMatch

The class declares variants that take a Locale and a Locale iterator.

#### LocaleMatcher::getBestMatchForListString

A separate member accepts a StringPiece locale list.

### Anchored code example

```
const Locale *LocaleMatcher::getBestMatch(const Locale &desiredLocale, UErrorCode &errorCode) const {
```

Source: `thirdparty/icu4c/common/localematcher.cpp` — LocaleMatcher::getBestMatch

Prerequisites: cpp-classes

Code examples: LocaleMatcher, Locale

Evidence:
- Code: `thirdparty/icu4c/common/localematcher.cpp` — LocaleMatcher::getBestMatch
- Code: `thirdparty/icu4c/common/unicode/localematcher.h` — LocaleMatcher
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## C++ plugin specialization

The pattern is repeated across scene, script, shader, resource, and version-control integrations.

C++ classes specialize EditorPlugin for feature-specific integrations; this requires C++ classes as the base abstraction mechanism.

### How it works

#### Scene plugin specialization

MeshLibraryEditorPlugin derives from EditorPlugin for mesh-library editing.

#### Script plugin specialization

EditorScriptPlugin derives from EditorPlugin for editor-script integration.

#### Shader plugin specialization

ShaderEditorPlugin derives from EditorPlugin for shader authoring integration.

### Anchored code example

```
class MeshLibraryEditorPlugin : public EditorPlugin {
```

Source: `editor/scene/3d/mesh_library_editor_plugin.h` — MeshLibraryEditorPlugin

Prerequisites: cpp-classes-and-inheritance

Evidence:
- Code: `editor/scene/3d/mesh_library_editor_plugin.h` — MeshLibraryEditorPlugin : public EditorPlugin
- Code: `editor/script/editor_script_plugin.h` — EditorScriptPlugin : public EditorPlugin
- Code: `editor/shader/shader_editor_plugin.h` — ShaderEditorPlugin : public EditorPlugin
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C-managed concurrent state

SDL wraps synchronization behind its own mutex and atomic APIs while its C code explicitly maintains the protected data.

C structs hold mutexes, atomic counts, and linked entries when SDL manages concurrent event and device state.

### How it works

#### SDL_EventQ in SDL_events.c

The queue combines SDL_Mutex, SDL_AtomicInt, and linked queue pointers.

#### SDL_StartEventLoop in SDL_events.c

Event-loop startup creates and locks the queue mutex before activating the queue and event-watch list.

### Anchored code example

```
SDL_Mutex *lock;

bool active;

SDL_AtomicInt count;
```

Source: `thirdparty/sdl/events/SDL_events.c` — SDL_EventQ

Prerequisites: c-structs-pointers

Code examples: SDL_EventEntry

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c` — SDL_EventQ
- Code: `thirdparty/sdl/events/SDL_events.c` — SDL_StartEventLoop
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## C explicit resource lifecycles

The Bzip2 adapter names initialization, reset, finalization, I/O, and close functions separately and wires them into FT_Stream.

C functions establish, reset, and release C struct state explicitly through paired lifecycle operations.

### How it works

#### Initialize then publish

FT_Stream_OpenBzip2 allocates a wrapper, initializes it, and stores it in stream->descriptor.pointer only after success.

#### Backward-seek reset

ft_bzip2_file_io resets decompression when a requested position precedes the current output position or reset is flagged.

#### Close owns wrapper cleanup

ft_bzip2_stream_close calls ft_bzip2_file_done, frees the wrapper, and clears descriptor.pointer.

### Anchored code example

```
if ( !FT_QNEW( zip ) )
{
  error = ft_bzip2_file_init( zip, stream, source );
  if ( error )
  {
    FT_FREE( zip );
    goto Exit;
  }

  stream->descriptor.pointer = zip;
}
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — FT_Stream_OpenBzip2

Prerequisites: c-structures-and-pointers

Code examples: FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_file_done
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_stream_close
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C function-pointer tables

SDL declares native-library symbols and device-driver operations as typed function pointers.

C function-pointer tables attach implementation callbacks to C struct state for dynamically selected system and device services.

### How it works

#### SDL_UDEV_Symbols in SDL_udev.h

The UDev adapter stores function pointers for libudev operations such as action and device-node queries.

#### hidapi_backend in SDL_hidapi.c

The HID layer defines backend operations behind function pointers and selects a backend per device.

### Anchored code example

```
const char *(*udev_device_get_action)(struct udev_device *);
const char *(*udev_device_get_devnode)(struct udev_device *);
```

Source: `thirdparty/sdl/core/linux/SDL_udev.h` — SDL_UDEV_Symbols

Prerequisites: c-structs-pointers

Code examples: SDL_hid_device_info

Evidence:
- Code: `thirdparty/sdl/core/linux/SDL_udev.h` — SDL_UDEV_Symbols
- Code: `thirdparty/sdl/hidapi/SDL_hidapi.c` — hidapi_backend
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## C stateful struct APIs

PSA, Mbed TLS TLS, and message-processing code all use this idiom.

C APIs expose mutable state through struct pointers, initialize it explicitly, and communicate failures through status-returning functions and caller-owned buffers.

### How it works

#### psa_reset_key_attributes

The function resets a caller-provided key-attribute object with memset.

#### mbedtls_ssl_session_save

The session serializer accepts a session pointer plus an output buffer, capacity, and output-length pointer.

#### mbedtls_mps_reader

The message-processing stack defines a reader structure whose implementation separates memory loads and stores from other operations.

### Anchored code example

```
void psa_reset_key_attributes(psa_key_attributes_t *attributes)
{
    memset(attributes, 0, sizeof(*attributes));
}
```

Source: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_client.c` — psa_reset_key_attributes

Code examples: psa_key_attributes_t, mbedtls_ssl_session

Evidence:
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_client.c` — psa_reset_key_attributes
- Code: `thirdparty/mbedtls/library/ssl_tls.c` — mbedtls_ssl_session_save; mbedtls_ssl_session_load
- Code: `thirdparty/mbedtls/library/mps_reader.c` — mbedtls_mps_reader implementation
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ copy-on-write container storage

Capacity growth, exact allocation, reallocation, and copy-to-new-buffer operations are explicitly represented in CowData.

The container layer uses C++ container classes to separate CowData buffer allocation and copying from Vector-style value interfaces.

### How it works

#### Buffer growth

CowData names growth and next-capacity calculations before allocation or reallocation.

#### Copy-to-new-buffer path

CowData has a copy-to-new-buffer operation for mutations that require separate storage.

#### Vector façade

Vector exposes data, iterators, and size operations over its underlying container storage.

### Anchored code example

```
const CowData prev_data;
```

Source: `core/templates/cowdata.h` — CowData::copy_from

Prerequisites: cpp-classes-inheritance

Code examples: Array, Dictionary

Evidence:
- Code: `core/templates/cowdata.h` — CowData
- Code: `core/templates/vector.h` — Vector
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C stateful streaming APIs

Brotli exposes decoder state through C interfaces while storing its operational fields in an internal structure.

C-facing code represents decoder progress as pointer-rich state and accepts caller-supplied allocation callbacks.

### How it works

#### Allocator callback default

BrotliDefaultAllocFunc accepts an opaque caller context and allocation size, then returns malloc-managed memory.

#### Stateful output retrieval

BrotliDecoderTakeOutput accepts a BrotliDecoderState pointer so output is obtained from maintained decoding state.

#### Internal state layout

BrotliDecoderStateStruct includes byte-pointer fields used by decoding internals.

### Anchored code example

```
void* BrotliDefaultAllocFunc(void* opaque, size_t size) {
  BROTLI_UNUSED(opaque);
  return malloc(size);
}
```

Source: `thirdparty/brotli/common/platform.c` — BrotliDefaultAllocFunc

Code examples: BrotliDecoderState

Evidence:
- Code: `thirdparty/brotli/common/platform.c` — BrotliDefaultAllocFunc
- Code: `thirdparty/brotli/dec/decode.c` — BrotliDecoderTakeOutput
- Code: `thirdparty/brotli/dec/state.h` — BrotliDecoderStateStruct
- External (official, verified): [ISO/IEC 9899:2018 — Information technology — Programming languages — C](https://www.iso.org/standard/74528.html), accessed 2026-07-16

## Callables and lambdas

The tests use typed parameters and returns, direct `call()`, captures, binding, and callable introspection.

Callable values can reference functions, constructors, built-ins, and lambdas; lambda bodies capture their enclosing context.

### How it works

#### Typed lambda invocation

Lambdas with zero, one, and typed return values are stored and invoked through `call()`.

#### Captured state

A lambda reads a surrounding local value and a member value when called later.

### Anchored code example

```
func test():
	var lambda_0 := func() -> void:
		print(0)
	lambda_0.call()
```

Source: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd` — test()

Prerequisites: function-contracts

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd` — test()
- Code: `modules/gdscript/tests/scripts/runtime/features/lambda_use_self.gd` — lambda1, lambda2, nested, lambda3, and lambda4
- Code: `modules/gdscript/tests/scripts/runtime/features/ctor_as_callable.gd` — inner_ctor and native_ctor
- External (official, verified): [Callable class reference](https://docs.godotengine.org/en/stable/classes/class_callable.html), accessed 2026-07-16

## C++ engine binding macros

These macros are a Godot-specific layer over C++ declarations used throughout editor extension points.

Godot's C++ classes declare engine binding and script-overridable callbacks through macros such as GDCLASS and GDVIRTUAL.

### How it works

#### Class declaration marker

EditorPlugin invokes GDCLASS immediately inside its class declaration.

#### Script-visible importer hooks

EditorSceneFormatImporter declares required virtual extension and import callbacks with GDVIRTUAL macros.

#### Inspector parser hooks

EditorInspectorPlugin declares script-overridable handling and parsing callbacks with GDVIRTUAL macros.

### Anchored code example

```
GDCLASS(EditorPlugin, Node);
```

Source: `editor/plugins/editor_plugin.h` — EditorPlugin

Prerequisites: cpp-classes-inheritance

Evidence:
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin
- Code: `editor/import/3d/resource_importer_scene.h` — EditorSceneFormatImporter
- Code: `editor/inspector/editor_inspector.h` — EditorInspectorPlugin
- External (primary, verified): [C++ Working Draft — Preprocessing directives](https://eel.is/c++draft/cpp), accessed 2026-07-16

## C++ interface polymorphism

The public extension classes mirror native polymorphic interfaces and distinguish required callbacks from optional ones.

C++ base classes and virtual callbacks define extension contracts whose subclasses provide physics, rendering, stream, and text implementations.

### How it works

#### RenderingDevice inheritance

RenderingDevice derives from RenderingDeviceCommons and uses Godot class metadata macros to participate in the native object model.

#### Physics virtual callbacks

PhysicsServer2DExtension declares virtual required operations such as _area_add_shape for external physics implementations.

#### Rendering buffer configuration

RenderSceneBuffersExtension exposes _configure so an extension handles viewport reconfiguration.

### Anchored code example

```
class RenderingDevice : public RenderingDeviceCommons {
	GDCLASS(RenderingDevice, Object)
	_THREAD_SAFE_CLASS_
```

Source: `servers/rendering/rendering_device.h` — RenderingDevice

Code examples: RenderSceneBuffersConfiguration, RID

Evidence:
- Code: `servers/rendering/rendering_device.h` — RenderingDevice
- Code: `doc/classes/PhysicsServer2DExtension.xml` — PhysicsServer2DExtension._area_add_shape
- Code: `doc/classes/RenderSceneBuffersExtension.xml` — RenderSceneBuffersExtension._configure
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## JavaScript browser runtime libraries

The Web platform places browser-facing implementation in JavaScript files while C++ classes call their exported bridge functions.

JavaScript libraries implement engine startup, preloading, runtime heap access, browser display and input callbacks, audio worklets, fetch, MIDI, and JavaScript object proxies.

### How it works

#### Engine

The engine library creates an Engine closure and uses Preloader and configuration objects to initialize the Web build.

#### GodotRuntime

Runtime helpers allocate strings and access typed heap views used to exchange WebAssembly memory with JavaScript.

#### GodotJSWrapper

The object bridge maps proxied JavaScript values and callbacks to IDs and conversion callbacks.

#### GodotProcessor

The audio worklet implementation consumes and produces audio through a RingBuffer and AudioWorkletProcessor.

### Anchored code example

```
const Engine = (function () {
```

Source: `platform/web/js/engine/engine.js` — Engine

Code examples: JavaScriptObjectImpl

Evidence:
- Code: `platform/web/js/engine/engine.js` — Engine
- Code: `platform/web/js/libs/library_godot_runtime.js` — GodotRuntime
- Code: `platform/web/js/libs/library_godot_javascript_singleton.js` — GodotJSWrapper
- Code: `platform/web/js/libs/audio.worklet.js` — GodotProcessor
- External (primary, verified): [ECMAScript Language Specification](https://tc39.es/ecma262/), accessed 2026-07-16

## Node paths and indexed access

Completion expectations vary with the source expression and with the node and script information available from a scene.

The fixtures use `$`, `%`, and indexed self access to resolve Node members inside a scene-aware test context.

### How it works

#### Indexed self access

A `Node` name is read and written through `self["name"]`.

#### Scene-informed completion

Scene-backed completion cases distinguish native, scripted, uniquely named, broad, and incompatible node references.

### Anchored code example

```
name = "Node"
	print(self["name"])
	self["name"] = "Changed"
```

Source: `modules/gdscript/tests/scripts/analyzer/features/subscript_self.gd` — test()

Code examples: Scene Fixture, Completion Test Configuration

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/subscript_self.gd` — test()
- Code: `modules/gdscript/tests/scripts/completion/get_node/literal_scene/dollar_unique.cfg` — [input] and [output]
- Code: `modules/gdscript/tests/scripts/completion/get_node/local_typehint_scene_incompatible/native_local_typehint_scene.cfg` — [output]
- External (official, unverified (source anchor not found)): [Node class reference](https://docs.godotengine.org/en/stable/classes/class_node.html), accessed 2026-07-15

## Classes, inheritance, and lookup

Fixtures test nested and external resolution separately from runtime inheritance behavior.

Classes supply nested types, inheritance, `super`, and lexical member lookup; preloaded classes also participate in lookup.

### How it works

#### Inner construction

Nested classes construct both themselves and sibling inner classes through their class names.

#### Inherited nested lookup

A nested class inherits from an outer class and resolves a constant through the inherited base.

### Anchored code example

```
class A:
	func a():
		return A.new()
	func b():
		return B.new()
```

Source: `modules/gdscript/tests/scripts/analyzer/features/inner_class_access_from_inside.gd` — A.a() and A.b()

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/inner_class_access_from_inside.gd` — A and B
- Code: `modules/gdscript/tests/scripts/analyzer/features/lookup_class.gd` — Y.B.check()
- Code: `modules/gdscript/tests/scripts/parser/features/super.gd` — SayAnotherThing
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C++ virtual interfaces

This is the main extension idiom in the inspected editor code.

Godot uses polymorphic base interfaces so importers, inspector plug-ins, preview generators, and editor plug-ins can supply type-specific behavior through virtual methods.

### How it works

#### Inspector eligibility

EditorInspectorPlugin exposes can_handle and parsing hooks so an implementation can select and build controls for an edited Object.

#### Scene-format operation

EditorSceneFormatImporter declares virtual extension discovery, scene import, option collection, and visibility methods.

#### Editor lifecycle hooks

EditorPlugin declares virtual methods for editing objects, reacting to UI input, saving state, building, and running scenes.

### Anchored code example

```
virtual bool can_handle(Object *p_object);
```

Source: `editor/inspector/editor_inspector.h` — EditorInspectorPlugin::can_handle

Prerequisites: cpp-classes-inheritance

Evidence:
- Code: `editor/inspector/editor_inspector.h` — EditorInspectorPlugin
- Code: `editor/import/3d/resource_importer_scene.h` — EditorSceneFormatImporter
- Code: `editor/plugins/editor_plugin.h` — EditorPlugin
- External (primary, verified): [C++ Working Draft — Virtual functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-16

## C++ virtual dispatch

This is used where ICU needs an interface with multiple concrete implementations.

Virtual dispatch uses derived classes to substitute behavior, as locale iterators override next and ICU break engines derive from a common engine type.

### How it works

#### Locale::RangeIterator::next

The iterator provides its concrete next operation with the override specifier.

#### LanguageBreakEngine

The base engine type supports specialized dictionary, unhandled, wrapper, and LSTM break-engine classes.

### Anchored code example

```
const Locale &next() override { return *it_++; }
```

Source: `thirdparty/icu4c/common/unicode/locid.h` — Locale::RangeIterator::next

Prerequisites: cpp-classes

Code examples: LocaleMatcher, RuleBasedBreakIterator

Evidence:
- Code: `thirdparty/icu4c/common/unicode/locid.h` — Locale::RangeIterator::next
- Code: `thirdparty/icu4c/common/brkeng.h` — LanguageBreakEngine
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## JavaScript browser FFI

The bridges use runtime allocation and parsing helpers, ID mapping, browser event callbacks, and Web API object methods.

JavaScript bridge libraries marshal browser objects, strings, buffers, and callbacks between web APIs and C++ platform code.

### How it works

#### RTC data-channel creation

The WebRTC bridge parses a configuration object and invokes createDataChannel on the browser peer connection.

#### WebSocket event callbacks

The WebSocket bridge retrieves runtime functions and binds open, message, error, and close callbacks.

#### WebXR session and input bridge

The WebXR bridge parses session options, allocates callback strings, and maintains browser-side frame and input-source state.

### Anchored code example

```
const channel = ref.createDataChannel(label, config);
```

Source: `modules/webrtc/library_godot_webrtc.js` — GodotRTCPeerConnection

Code examples: WebRTCPeerConnection, WebRTCDataChannel, WebXRInterfaceJS

Evidence:
- Code: `modules/webrtc/library_godot_webrtc.js` — GodotRTCPeerConnection
- Code: `modules/websocket/library_godot_websocket.js` — GodotWebSocket
- Code: `modules/webxr/native/library_godot_webxr.js` — GodotWebXR
- External (primary, verified): [ECMAScript Language Specification](https://tc39.es/ecma262/multipage/ecmascript-language-functions-and-classes.html#sec-arrow-function-definitions), accessed 2026-07-16

## Signals and await

Fixtures distinguish meaningful awaits from redundant ones and cover signal payload shapes.

Signal values and callable functions support waiting for emissions or coroutine completion with `await`.

### How it works

#### Signal identity

Signals declared at outer, nested, and inherited classes are compared for identity.

#### Await payload shape

Awaiting signals with zero, one, and two parameters produces distinct returned values in the test output.

### Anchored code example

```
func await_one_parameter():
	var result = await one_parameter
```

Source: `modules/gdscript/tests/scripts/runtime/features/await_signal_with_parameters.gd` — await_one_parameter()

Prerequisites: callables-and-lambdas

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/runtime/features/await_signal_with_parameters.gd` — await_one_parameter()
- Code: `modules/gdscript/tests/scripts/analyzer/features/lookup_signal.gd` — get_signal()
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/redundant_await.out` — REDUNDANT_AWAIT
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

## C++ types, encapsulation, and inheritance

The partition uses classes as stable subsystem boundaries and exposes operations through member functions.

This C++ code models decoders, geometry engines, and allocators as named types with inheritance and encapsulated state.

### How it works

#### Stream specialization

jpeg_decoder_file_stream and jpeg_decoder_mem_stream derive from jpeg_decoder_stream, separating file and memory input implementations behind a common decoder-stream type.

#### Geometry engine specialization

Clipper64 and ClipperD derive from ClipperBase while retaining integer and scaled floating-point result construction.

#### Encapsulated transcoder state

basisu_transcoder stores its low-level decoder members privately and exposes accessor methods for them.

### Anchored code example

```
class png_memory_file : public png_file
```

Source: `thirdparty/basis_universal/encoder/pvpngreader.cpp` — png_memory_file

Code examples: Ktx2TranscoderState, PolyPathD

Evidence:
- Code: `thirdparty/basis_universal/encoder/pvpngreader.cpp` — png_memory_file
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — Clipper64
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — basisu_transcoder
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

## C++ byte-exact binary parsing

The texture containers deliberately operate on byte-addressed buffers and packed integer fields.

This C++ code reads byte streams through packed fields, pointer casts, and bit readers; this is necessary to map serialized texture headers to in-memory values.

### How it works

#### Header reinterpretation

Basis transcoding casts an input byte pointer to basis_file_header before obtaining offsets and slice descriptors.

#### Packed KTX2 fields

ktx2_header and ktx2_level_index use packed_uint fields for on-disk values, while pragma packing surrounds the KTX2 layout declarations.

#### Bitwise decoder cursor

bitwise_decoder maintains start, current, and end byte pointers while providing bit-reading operations.

### Anchored code example

```
const basis_file_header* pHeader = reinterpret_cast<const basis_file_header*>(pData);
```

Source: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_transcoder::get_file_info

Prerequisites: cpp-abstraction-and-polymorphism

Code examples: BasisFileHeader, BasisSliceDescriptor, Ktx2Header, Ktx2LevelIndex

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp` — basisu_transcoder::get_file_info
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder_internal.h` — bitwise_decoder
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

## C++ templates and generic containers

Basis containers, Clipper points and paths, and Embree math types use type parameters rather than duplicate implementations per element type.

This C++ code parameterizes collection and math utilities by type, so their typed operations can be reused.

### How it works

#### Basis vector container

basisu::vector is a templated collection type that derives relational operations from its element-specialized form.

#### Generic geometry points

Clipper's Point<T> supports coordinate representations selected by the template parameter.

#### Generic SIMD-adjacent math

Embree's Vec3<T> and BBox<T> operate on a supplied scalar or vector-like element type.

### Anchored code example

```
class vector : public rel_ops< vector<T> >
```

Source: `thirdparty/basis_universal/transcoder/basisu_containers.h` — basisu::vector

Prerequisites: cpp-abstraction-and-polymorphism

Code examples: Ktx2LevelIndex, PolyPathD

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_containers.h` — basisu::vector
- Code: `thirdparty/clipper2/include/clipper2/clipper.core.h` — Clipper2Lib::Point
- Code: `thirdparty/embree/common/math/vec3.h` — embree::Vec3
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16

## C++ class hierarchies and reference bases

The code combines reusable base classes with specialized derived classes rather than representing all runtime objects as unrelated records.

C++ class inheritance connects shared reference ownership to acceleration, scene, geometry, builder, and scheduler abstractions.

### How it works

#### Reference-count base

RefCount underlies retained objects; AccelData, Builder, and other abstractions inherit from it.

#### Acceleration hierarchy

AccelData is the base for acceleration data, while Accel and AccelN organize intersector and aggregate-acceleration behavior.

#### Geometry hierarchy

Geometry is a common base, and TriangleMesh, GridMesh, Instance, and other concrete forms derive from it.

### Anchored code example

```
class AccelData : public RefCount
```

Source: `thirdparty/embree/kernels/common/accel.h` — AccelData

Code examples: RefCount, AccelData, Geometry, Builder

Evidence:
- Code: `thirdparty/embree/common/sys/ref.h` — RefCount
- Code: `thirdparty/embree/kernels/common/accel.h` — AccelData
- Code: `thirdparty/embree/kernels/common/geometry.h` — Geometry
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## Annotations, static state, and lifecycle

The suite covers annotation validity, tool-mode requirements, static initialization, and ready-time values.

Annotations modify class declarations and members, while static state and `@onready` participate in initialization and lifecycle behavior.

### How it works

#### Ready-time member

An inner `Node` subclass declares an `@onready` member even when the outer script extends `RefCounted`.

#### Static initialization

Static-variable fixtures observe static initialization and state across loaded scripts and nested classes.

### Anchored code example

```
class Inner extends Node:
	@onready var okay = 0
```

Source: `modules/gdscript/tests/scripts/analyzer/features/onready_on_inner_class_with_non_node_outer.gd` — Inner

Prerequisites: classes-inheritance-and-lookup

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/onready_on_inner_class_with_non_node_outer.gd` — Inner
- Code: `modules/gdscript/tests/scripts/runtime/features/static_variables.gd` — Inner and InnerInner
- Code: `modules/gdscript/tests/scripts/parser/errors/export_tool_button_requires_tool_mode.out` — expected analyzer error
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C++ C-linkage adapters

Graphite exposes C-callable functions that delegate to C++ objects such as CharInfo, Face, Segment, and Slot.

C++ classes are wrapped behind C linkage functions and derived opaque API types, so C callers use handles rather than C++ class declarations.

### How it works

#### CharInfo C entry points

The gr_cinfo_* functions have C linkage, validate the handle with assert, and delegate to CharInfo methods.

#### Opaque derived handles

Types such as gr_face, gr_font, gr_segment, and gr_slot derive from their C++ implementation types and are the types accepted by the C API.

### Anchored code example

```
extern "C"
{

unsigned int gr_cinfo_unicode_char(const gr_char_info* p/*not NULL*/)
{
    assert(p);
    return p->unicodeChar();
}
```

Source: `thirdparty/graphite/src/gr_char_info.cpp` — gr_cinfo_unicode_char

Prerequisites: cpp-classes-and-inheritance

Code examples: Face, Font, Segment, Slot

Evidence:
- Code: `thirdparty/graphite/src/gr_char_info.cpp` — extern "C" and gr_cinfo_unicode_char
- Code: `thirdparty/graphite/src/inc/CharInfo.h` — gr_char_info
- Code: `thirdparty/graphite/src/inc/Face.h` — gr_face
- External (primary, unverified (source anchor not found)): [C++ working draft: Linkage specifications](https://eel.is/c++draft/dcl.link), accessed 2026-07-15

## C++ class inheritance

The source uses base-class declarations to organize runtime, GUI, and asset types.

C++ class declarations define the repository's object families through public base classes, including Node-derived runtime objects, Resource-derived assets, and Viewport-derived windows.

### How it works

#### Node root

Node is declared as an Object subclass, establishing the root of the scene runtime family.

#### Viewport specialization

Window publicly derives from Viewport, so window behavior is represented as a viewport specialization.

#### Resource specialization

PackedScene publicly derives from Resource, placing serialized scenes in the reusable asset family.

### Anchored code example

```
class Node : public Object {
```

Source: `scene/main/node.h` — class Node

Code examples: Node, PackedScene

Evidence:
- Code: `scene/main/node.h` — class Node : public Object
- Code: `scene/main/window.h` — class Window : public Viewport
- Code: `scene/resources/packed_scene.h` — class PackedScene : public Resource
- External (primary, verified): [Working Draft, Programming Languages — C++: Derived classes](https://eel.is/c++draft/class.derived), accessed 2026-07-16

## C++ scoped lock guards

Synchronization wrappers are declared in the OS layer and reused by object and reflection code.

The implementation uses a scoped C++ MutexLock object and class-specific lock helpers around shared runtime state.

### How it works

#### Mutex lock wrapper

MutexLock is a nodiscard class in the OS synchronization header.

#### ClassDB lock types

ClassDB declares Locker and Lock helper classes for its registry-related work.

#### Safe binary mutex specialization

MutexLock is specialized for SafeBinaryMutex tags.

### Anchored code example

```
class [[nodiscard]] MutexLock {
```

Source: `core/os/mutex.h` — MutexLock

Prerequisites: cpp-classes-inheritance

Evidence:
- Code: `core/os/mutex.h` — MutexLock
- Code: `core/object/class_db.h` — ClassDB::Lock
- Code: `core/os/safe_binary_mutex.h` — MutexLock<SafeBinaryMutex<Tag>>
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ templates and non-owning views

The rendering code combines engine containers with pointer-and-length style views for recorded GPU work.

Template-based containers and views expose typed data without copying it, including VectorView access to command and resource arrays.

### How it works

#### Typed view access

VectorView stores a templated pointer and returns a const reference from indexed access.

#### Typed conversion helpers

VariantConverterStd140 specializations and conversion helpers translate variant values and arrays into GPU uniform layouts.

Code examples: RecordedCommand

Evidence:
- Code: `servers/rendering/rendering_device_commons.h` — VectorView
- Code: `servers/rendering/storage/variant_converters.h` — VariantConverterStd140
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4928.pdf), accessed 2026-07-15

## Scoped enums and bit flags

The code uses compact underlying integer types for many externally meaningful simulation modes.

Scoped enum class values model motion, body, collision, update-error, and configuration states, while selected enums provide bitwise combinations.

### How it works

#### Motion state

EMotionType distinguishes Static, Kinematic, and Dynamic behavior.

#### Degrees of freedom

EAllowedDOFs encodes translational and rotational permissions as bit values and defines combinations such as Plane2D.

### Anchored code example

```
enum class EMotionType : uint8
{
	Static,						///< Non movable
	Kinematic,					///< Movable using velocities only, does not respond to forces
	Dynamic,					///< Responds to forces as a normal physics object
};
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h` — EMotionType

Code examples: Body, ConstraintSettings

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionType.h` — EMotionType
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/AllowedDOFs.h` — EAllowedDOFs and operator |
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## C++ static casts across track types

Animation evaluation repeats this pattern for several concrete track classes.

The implementation uses static_cast to interpret a base Track pointer as a derived Track object after a track kind has been selected.

### How it works

#### Value track selection

Animation code casts a selected Track pointer to ValueTrack before accessing value-track keys.

#### Other typed tracks

The same evaluation area casts Track pointers to BezierTrack, AudioTrack, and AnimationTrack variants.

### Anchored code example

```
const ValueTrack *vt = static_cast<const ValueTrack *>(tracks[track]);
```

Source: `scene/resources/animation.cpp` — Animation track evaluation

Prerequisites: cpp-object-hierarchies

Code examples: Animation

Evidence:
- Code: `scene/resources/animation.cpp` — static_cast<const ValueTrack *>(tracks[track])
- Code: `scene/resources/animation.cpp` — static_cast selections for BezierTrack, AudioTrack, and AnimationTrack
- External (primary, verified): [Working Draft, Programming Languages — C++: Static cast](https://eel.is/c++draft/expr.static.cast), accessed 2026-07-16

## C++ static generated data

The repository contains generated HarfBuzz UCD tables and generated ICU normalization, bidi, case, and character-property tables.

Static const arrays and values embed generated Unicode and normalization data directly in translation units.

### How it works

#### norm2_nfc_data_formatVersion

Normalization data records format and data versions as static constants.

#### _hb_ucd_sc_map

HarfBuzz embeds a generated script map as a static array.

### Anchored code example

```
static const UVersionInfo norm2_nfc_data_formatVersion={5,0,0,0};
static const UVersionInfo norm2_nfc_data_dataVersion={0x11,0,0,0};
```

Source: `thirdparty/icu4c/common/norm2_nfc_data.h` — norm2_nfc_data_formatVersion

Code examples: UCPTrie

Evidence:
- Code: `thirdparty/icu4c/common/norm2_nfc_data.h` — norm2_nfc_data_formatVersion
- Code: `thirdparty/harfbuzz/src/hb-ucd-table.hh` — _hb_ucd_sc_map
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## C++ templates

The code uses template parameters for both general utility code and statically selected font-table implementations.

Templates parameterize data structures and algorithms, including string code-point appenders and typed OpenType table subset calls.

### How it works

#### appendCodePoint

A StringClass template selects a code-unit representation while a boolean template parameter controls validation.

#### _hb_subset_table

Table-tag cases instantiate _hb_subset_table with a concrete OpenType table type.

### Anchored code example

```
template<typename StringClass, bool validate>
inline StringClass &appendCodePoint(StringClass &s, uint32_t c) {
```

Source: `thirdparty/icu4c/common/unicode/utfstring.h` — utfstring::prv::appendCodePoint

Code examples: UCPTrie, hb_subset_plan_t

Evidence:
- Code: `thirdparty/icu4c/common/unicode/utfstring.h` — utfstring::prv::appendCodePoint
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## C++ classes, inheritance, and polymorphic interfaces

The implementation favors abstract base types with specialized derived classes.

C++ classes and inheritance define extensible interfaces for shapes, constraints, collision queries, job systems, renderers, and body access.

### How it works

#### Shape hierarchy

Shape derives from reference-counting and non-copyable bases, while ConvexShape, CompoundShape, mesh, terrain, and decorated types provide specialized collision behavior.

#### Read and write body locks

BodyLockRead and BodyLockWrite specialize a common lock base with const and mutable Body access.

### Anchored code example

```
class JPH_EXPORT Shape : public RefTarget<Shape>, public NonCopyable
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h` — Shape

Code examples: Shape, Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h` — Shape
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/BodyLock.h` — BodyLockBase, BodyLockRead, BodyLockWrite
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## Properties and accessors

Tests cover direct assignment, backing variables, getter/setter syntax, and compound access chains.

Class properties route reads and writes through inline or named accessors, including static and chained access.

### How it works

#### Named accessor binding

A property delegates its getter and setter to separately declared methods.

#### Inline backing property

An inline property returns and updates a separate backing field.

### Anchored code example

```
var prop:
	get = get_prop, set = set_prop

func get_prop():
	return _prop
```

Source: `modules/gdscript/tests/scripts/analyzer/features/property_functions.gd` — prop and get_prop()

Prerequisites: classes-inheritance-and-lookup

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/property_functions.gd` — prop
- Code: `modules/gdscript/tests/scripts/analyzer/features/property_inline.gd` — prop4
- Code: `modules/gdscript/tests/scripts/runtime/features/setter_chain_shared_types.gd` — prop1, prop2, prop3, and prop4
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C++ frame-time arithmetic

`MainTimerSync` contains the timing policy; `DeltaSmoother` is a separate timing-related class in the same subsystem.

C++ frame timing turns accumulated time and a physics tick rate into bounded physics-step counts for frame timing.

### How it works

#### MainTimerSync::advance

The implementation computes a maximum possible step count with `std::floor`, accumulated time, the physics tick rate, and a jitter-fix value.

#### MainTimerSync

`MainTimerSync` is the class that owns the timing synchronization API.

#### DeltaSmoother

`DeltaSmoother` is declared beside `MainTimerSync`, separating delta smoothing from the timer-sync class.

### Anchored code example

```
const int max_possible_steps = std::floor((time_accum)*p_physics_ticks_per_second + get_physics_jitter_fix());
```

Source: `main/main_timer_sync.cpp` — MainTimerSync::advance

Code examples: MainFrameTime

Evidence:
- Code: `main/main_timer_sync.h` — MainTimerSync
- Code: `main/main_timer_sync.h` — DeltaSmoother
- Code: `main/main_timer_sync.cpp` — MainTimerSync::advance
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

## C++ inheritance interfaces

The partition consistently uses derived classes to plug implementations into engine-facing abstractions.

C++ classes specialize engine and OpenXR base interfaces through public inheritance.

### How it works

#### Spatial extension specialization

OpenXRSpatialEntityExtension derives from OpenXRExtensionWrapper, placing spatial behavior behind the wrapper base type.

#### Image-loader specialization

ImageLoaderSVG and ImageLoaderTGA derive from ImageFormatLoader, expressing the image-format adapter contract in their class declarations.

#### Text-server specialization

TextServerAdvanced and TextServerFallback each derive from TextServerExtension, providing alternate implementations behind one engine-facing type.

### Anchored code example

```
class OpenXRSpatialEntityExtension : public OpenXRExtensionWrapper {
```

Source: `modules/openxr/extensions/spatial_entities/openxr_spatial_entity_extension.h` — OpenXRSpatialEntityExtension declaration

Code examples: OpenXRStructureBase, ShapedTextDataAdvanced

Evidence:
- Code: `modules/openxr/extensions/spatial_entities/openxr_spatial_entity_extension.h` — class OpenXRSpatialEntityExtension : public OpenXRExtensionWrapper
- Code: `modules/svg/image_loader_svg.h` — class ImageLoaderSVG : public ImageFormatLoader
- Code: `modules/text_server_adv/text_server_adv.h` — class TextServerAdvanced : public TextServerExtension
- External (primary, verified): [Working Draft, Programming Languages — C++](https://eel.is/c++draft/class.derived), accessed 2026-07-16

## C# partial types and source-generator tests

Partial declarations let generated files add members to a type declared in the sample or test source.

The .NET SDK test projects define partial C# types and verify generators that emit Godot-facing signal, property, method, serialization, and script-path code.

### How it works

#### Signal-bearing partial type

The EventSignals sample declares a partial GodotObject class with a Signal-attributed delegate.

#### Generator verification

ScriptPropertiesGeneratorTests invokes a CSharpSourceGeneratorVerifier for exported fields and properties.

#### Generated serialization implementation

The checked-in generated serialization files override SaveGodotObjectData and RestoreGodotObjectData.

### Anchored code example

```
public partial class EventSignals : GodotObject
{
    [Signal]
    public delegate void MySignalEventHandler(string str, int num);
}
```

Source: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Sample/EventSignals.cs` — EventSignals

Code examples: CSharpScript

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Sample/EventSignals.cs` — EventSignals
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/ScriptPropertiesGeneratorTests.cs` — ScriptPropertiesGeneratorTests
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators.Tests/TestData/GeneratedSources/ScriptBoilerplate_ScriptSerialization.generated.cs` — SaveGodotObjectData and RestoreGodotObjectData
- External (official, verified): [Partial type - C# reference](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/partial-type), accessed 2026-07-16

## Kotlin runtime and service coordination

Kotlin is used for the main `Godot` host, activities, services, remote fragments, storage handlers, and editor code.

Kotlin coordinates the Android platform host through nullable runtime state, Android host construction, and integer-coded remote-service messages.

### How it works

#### Godot

The main Kotlin host is a class with a private constructor that receives an Android `Context`.

#### GodotService

The service source defines constants for initialization, start, stop, destroy, engine error, status updates, and restart requests.

#### RemoteGodotFragment

The remote fragment imports Android binding and Messenger types and refers to the `GodotService` engine-status and engine-error types.

### Anchored code example

```
class Godot private constructor(val context: Context) {
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt` — Godot

Code examples: Callable, GodotService.EngineStatus

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/Godot.kt` — Godot
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt` — MSG_INIT_ENGINE
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/RemoteGodotFragment.kt` — RemoteGodotFragment
- External (official, unverified (source anchor not found)): [Classes | Kotlin Documentation](https://kotlinlang.org/docs/classes.html), accessed 2026-07-15

## C++ inline native wrappers

Metal-cpp declares lightweight C++ classes and forwards calls through typed `Object::sendMessage` invocations.

C++ wrapper classes use inheritance and inline member functions to expose native framework methods.

### How it works

#### MTL4::Archive inheritance

`MTL4::Archive` derives from `NS::Referencing<Archive>`, establishing the wrapper's framework-object base.

#### MTL4::Archive::newBinaryFunction

The inline method binds a typed return value, selector, descriptor argument, and error out-parameter in one forwarding call.

#### MTL4::BinaryFunction::name

A similarly inline wrapper returns an `NS::String*` by sending the `name` selector.

### Anchored code example

```
_MTL_INLINE MTL4::BinaryFunction* MTL4::Archive::newBinaryFunction(const MTL4::BinaryFunctionDescriptor* descriptor, NS::Error** error)
{
    return Object::sendMessage<MTL4::BinaryFunction*>(this, _MTL_PRIVATE_SEL(newBinaryFunctionWithDescriptor_error_), descriptor, error);
}
```

Source: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp` — MTL4::Archive::newBinaryFunction

Code examples: MTL4::Archive, MTL4::BinaryFunction

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTL4Archive.hpp` — MTL4::Archive
- Code: `thirdparty/metal-cpp/Metal/MTL4BinaryFunction.hpp` — MTL4::BinaryFunction::name
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C# partial types and source generation

Godot's analyzers and generators inspect C# syntax and semantic symbols, report invalid declarations, and add generated source to the compilation.

Compilation produces attribute-bearing partial declarations for script paths and assembly script types.

### How it works

#### Script path generation

ScriptPathAttributeGenerator.Execute discovers Godot script classes, rejects non-partial classes, emits ScriptPathAttribute declarations, and emits an AssemblyHasScriptsAttribute source file.

#### Member metadata generation

ScriptPropertiesGenerator, ScriptMethodsGenerator, and ScriptSignalsGenerator generate bridge-facing metadata from compatible members and delegates.

#### Unmanaged callback generation

UnmanagedCallbacksGenerator examines partial method definitions and emits callback trampolines and matching function-structure members.

Prerequisites: csharp-attributes-reflection

Code examples: ScriptPathAttribute, AssemblyHasScriptsAttribute, ManagedCallbacks

Evidence:
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs` — ScriptPathAttributeGenerator.Execute
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/Common.cs` — Common.ClassPartialModifierRule
- Code: `modules/mono/glue/GodotSharp/Godot.SourceGenerators.Internal/UnmanagedCallbacksGenerator.cs` — UnmanagedCallbacksGenerator
- External (official, unverified (source anchor not found)): [Partial type (C# Reference)](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/partial-type), accessed 2026-07-15

## C++ mapping of C API types

The generated mapping provides a type-level bridge between the C Vulkan ABI names and the C++ API layer.

`CppType<Vk...>` specializations associate C Vulkan type names with C++ wrapper types.

### How it works

#### Wrapper-type association

A `CppType` specialization is emitted for each mapped Vulkan C record type.

#### Structure-type association

Many record mappings are accompanied by a `CppType<StructureType, ...>` specialization keyed by the Vulkan structure-type enumerator.

### Anchored code example

```
struct CppType<VkGraphicsPipelineCreateInfo>
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — CppType<VkGraphicsPipelineCreateInfo>

Prerequisites: cpp-typed-api-records

Code examples: GraphicsPipelineCreateInfo

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — CppType<VkGraphicsPipelineCreateInfo>
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp` — CppType<StructureType, StructureType::eGraphicsPipelineCreateInfo>
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ class inheritance for backend contracts

The code uses public derived classes to specialize a common engine-facing interface.

C++ class inheritance expresses shared renderer and server contracts and allows concrete dummy, extension, wrapper, clustered, and mobile implementations.

### How it works

#### Physics backend substitution

PhysicsServer3DDummy, PhysicsServer3DExtension, and PhysicsServer3DWrapMT derive from PhysicsServer3D.

#### Renderer specialization

RenderForwardClustered and RenderForwardMobile derive from RendererSceneRenderRD.

#### Text and XR extension surfaces

TextServerExtension derives from TextServer and XRInterfaceExtension derives from XRInterface.

### Anchored code example

```
class PhysicsServer3DDummy : public PhysicsServer3D {
```

Source: `servers/physics_3d/physics_server_3d_dummy.h` — PhysicsServer3DDummy

Evidence:
- Code: `servers/physics_3d/physics_server_3d_dummy.h` — PhysicsServer3DDummy
- Code: `servers/rendering/renderer_rd/forward_clustered/render_forward_clustered.h` — RenderForwardClustered
- Code: `servers/rendering/renderer_rd/forward_mobile/render_forward_mobile.h` — RenderForwardMobile
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4928.pdf), accessed 2026-07-15

## C# unsafe interop and function pointers

GodotSharp validates unmanaged table size before reinterpreting its address, marshals engine values through native layouts, and uses generated callback tables for native-to-managed calls.

Unsafe C# stores pointers and function pointers in ABI structs and invokes generated partial native calls.

### How it works

#### Native function table initialization

NativeFuncs.Initialize verifies the unmanaged callback-structure size and stores a pointer-backed UnmanagedCallbacks value.

#### Managed callback ABI

ManagedCallbacks is sequentially laid out and contains unmanaged function pointers for signals, delegates, script management, and marshaling.

#### Variant argument access

NativeVariantPtrArgs is an unsafe ref struct that stores a godot_variant double pointer and exposes indexed references without enabling unsafe blocks in game projects.

Prerequisites: csharp-source-generation

Code examples: ManagedCallbacks, Variant

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeFuncs.cs` — NativeFuncs.Initialize
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs` — ManagedCallbacks
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeVariantPtrArgs.cs` — NativeVariantPtrArgs
- External (official, unverified (source anchor not found)): [Unsafe code, pointers to data, and function pointers - C# reference](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/unsafe-code), accessed 2026-07-15

## Python build code generation

This is build-time support rather than the engine's C++ runtime.

The Python build helper defines generation and build-entry functions for virtual-method source generation.

### How it works

#### Generation function

generate_version accepts argument-count and compatibility switches for emitted virtual-method variants.

#### Build entry function

run receives target, source, and environment parameters for build integration.

### Anchored code example

```
def generate_version(argcount, const=False, returns=False, required=False, compat=False):
```

Source: `core/object/make_virtuals.py` — generate_version

Evidence:
- Code: `core/object/make_virtuals.py` — generate_version
- Code: `core/object/make_virtuals.py` — run
- External (official, verified): [The Python Language Reference — Function definitions](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

## C function-pointer interfaces

The implementation assigns allocation and I/O functions into external library or stream interface records.

C function pointers attach behavior to struct-and-pointer handles without requiring a language-level object system.

### How it works

#### Bzip2 allocator hooks

ft_bzip2_file_init assigns ft_bzip2_alloc and ft_bzip2_free to bz_stream callbacks.

#### Stream callbacks

FT_Stream_OpenBzip2 installs ft_bzip2_stream_io and ft_bzip2_stream_close in the target stream.

#### Cache extension callbacks

The cache callback header declares new, free, weight, and comparison functions for image, small-bitmap, and generic cache nodes.

### Anchored code example

```
static unsigned long
ft_bzip2_stream_io( FT_Stream       stream,
                    unsigned long   offset,
                    unsigned char*  buffer,
                    unsigned long   count )
{
  FT_BZip2File  zip = (FT_BZip2File)stream->descriptor.pointer;


  return ft_bzip2_file_io( zip, offset, buffer, count );
}
```

Source: `thirdparty/freetype/src/bzip2/ftbzip2.c` — ft_bzip2_stream_io

Prerequisites: c-structures-and-pointers

Code examples: FT_Stream, FT_BZip2FileRec

Evidence:
- Code: `thirdparty/freetype/src/cache/ftccback.h` — ftc_inode_new
- External (primary, unverified (source anchor not found)): [N1570: Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ polymorphic platform adapters

The common export-platform base is extended by platform-specific exporters, including the Apple embedded exporter.

C++ class inheritance and virtual functions define the target-platform adapter contract.

### How it works

#### Common target contract

EditorExportPlatform declares the target-export interface that concrete platform exporters implement.

#### Apple specialization

EditorExportPlatformAppleEmbedded derives from EditorExportPlatform to provide Apple embedded behavior.

### Anchored code example

```
class EditorExportPlatformAppleEmbedded : public EditorExportPlatform {
```

Source: `editor/export/editor_export_platform_apple_embedded.h` — EditorExportPlatformAppleEmbedded

Code examples: EditorExportPlatform

Evidence:
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform
- Code: `editor/export/editor_export_platform_apple_embedded.h` — EditorExportPlatformAppleEmbedded
- External (primary, unverified (source anchor not found)): [Draft C++ Standard — Virtual functions](https://eel.is/c++draft/class.virtual), accessed 2026-07-15

## C++ inheritance, virtual interfaces, and Ref ownership

Core abstractions such as resources, input events, and resource format loaders are C++ class hierarchies.

C++ classes use inheritance, virtual methods, Ref ownership, and static registries to implement core services.

### How it works

#### ResourceFormatLoader::load

ResourceFormatLoader derives from RefCounted and declares virtual loading, recognition, dependency, and type-query operations that concrete format loaders override.

#### ResourceLoader::ThreadLoadTask

The loader keeps task, path, status, cache-mode, resource, dependency, and awaiting state in its nested threaded-load record.

#### InputEvent hierarchy

InputEvent derives from Resource, and InputEventFromWindow, InputEventWithModifiers, key, mouse, joypad, touch, gesture, MIDI, and action types extend that base.

### Anchored code example

```
class InputEvent : public Resource {
```

Source: `core/input/input_event.h` — InputEvent

Code examples: Resource, InputEvent, GDExtension

Evidence:
- Code: `core/io/resource_loader.h` — ResourceFormatLoader
- Code: `core/io/resource_loader.h` — ResourceLoader::ThreadLoadTask
- Code: `core/input/input_event.h` — InputEvent
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

## C++ pluggable allocation

The hook is explicit rather than relying on a global replacement of operator new.

meshoptimizer stores allocation and deallocation function pointers in static allocator storage so callers can replace the allocation policy.

### How it works

#### meshopt_Allocator::storage

storage returns the static allocation-function storage used by the allocator wrapper.

#### meshopt_setAllocator

The setter asserts both callbacks exist and stores them in the shared allocation storage.

#### meshopt_Allocator use sites

Mesh processing functions instantiate meshopt_Allocator before allocating temporary working memory.

### Anchored code example

```
void meshopt_setAllocator(void* (MESHOPTIMIZER_ALLOC_CALLCONV* allocate)(size_t), void (MESHOPTIMIZER_ALLOC_CALLCONV* deallocate)(void*))
```

Source: `thirdparty/meshoptimizer/allocator.cpp` — meshopt_setAllocator

Code examples: meshopt_Meshlet

Evidence:
- Code: `thirdparty/meshoptimizer/allocator.cpp` — meshopt_Allocator::storage; meshopt_setAllocator
- Code: `thirdparty/meshoptimizer/vfetchoptimizer.cpp` — meshopt_Allocator allocator
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4950](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ polymorphism and ownership

The OpenXR loader exposes abstract recorder types while factory functions return owning smart pointers.

C++ classes use public inheritance and `std::unique_ptr` to retain polymorphic loader services and dispatch state.

### How it works

#### LoaderLogRecorder

The recorder hierarchy supplies an interface for alternative logging destinations.

#### MakeStdErrLoaderLogRecorder

The factory returns `std::unique_ptr<LoaderLogRecorder>`, making ownership transfer explicit in the function type.

#### LoaderInstance::DispatchTable

The loader instance stores its generated dispatch table behind a unique pointer.

### Anchored code example

```
std::unique_ptr<LoaderLogRecorder> MakeStdErrLoaderLogRecorder(void* user_data);
```

Source: `thirdparty/openxr/src/loader/loader_logger_recorders.hpp` — MakeStdErrLoaderLogRecorder

Prerequisites: cpp-native-wrappers

Code examples: LoaderInstance, XrGeneratedDispatchTableCore

Evidence:
- Code: `thirdparty/openxr/src/loader/loader_logger.hpp` — LoaderLogRecorder
- Code: `thirdparty/openxr/src/loader/loader_instance.hpp` — LoaderInstance::DispatchTable
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## Java Android host APIs

Java classes supply the Android-facing half of rendering, input, plugins, and native-library integration.

Java uses class inheritance and interface implementation to attach GL/Vulkan views, plugin registration, native library calls, and input handling to the Android platform host.

### How it works

#### GodotGLRenderView

The GL render view extends `GLSurfaceView` and implements the engine `GodotRenderView` interface.

#### GodotVulkanRenderView

The Vulkan render view is a separate Java class implementing `GodotRenderView` through the Vulkan surface-view path.

#### GodotPluginRegistry

The plugin registry is a Java class in the host library and imports Android package metadata APIs and concurrent collections.

#### GodotLib

The Java native-library class imports rendering, storage, native-bridge, TTS, network, and callable types.

### Anchored code example

```
class GodotGLRenderView extends GLSurfaceView implements GodotRenderView {
```

Source: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java` — GodotGLRenderView

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotGLRenderView.java` — GodotGLRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotVulkanRenderView.java` — GodotVulkanRenderView
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/plugin/GodotPluginRegistry.java` — GodotPluginRegistry
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/GodotLib.java` — GodotLib
- External (official, verified): [The Java Language Specification, Chapter 8: Classes](https://docs.oracle.com/javase/specs/jls/se17/html/jls-8.html), accessed 2026-07-16

## Python build scripts

The inventory contains Python builders beside SCsub files that import them.

Python build scripts import builder modules and define builder functions used by SCons-facing scene-theme scripts.

### How it works

#### Builder-module import

The icon SCsub script imports default_theme_icons_builders.

#### Builder function definition

The default theme builder module defines make_fonts_header.

### Anchored code example

```
import default_theme_builders
```

Source: `scene/theme/SCsub` — module import

Evidence:
- Code: `scene/theme/SCsub` — import default_theme_builders
- Code: `scene/theme/default_theme_builders.py` — make_fonts_header
- Code: `scene/theme/icons/default_theme_icons_builders.py` — make_default_theme_icons_action
- External (official, verified): [The Python Language Reference: The import system](https://docs.python.org/3/reference/import.html), accessed 2026-07-16

## C function-pointer callbacks

The ENet callback record is initialized with standard allocation functions and can be replaced through its initialization API.

C function-pointer fields let ENet obtain allocator, deallocator, and out-of-memory handlers through a cross-cutting callbacks record.

### How it works

#### Callback record definition

ENetCallbacks stores three callable function-pointer fields for allocation, deallocation, and allocation failure handling.

#### Runtime callback installation

enet_initialize_with_callbacks validates supplied allocation callbacks and assigns them to the process-level callbacks record.

### Anchored code example

```
typedef struct _ENetCallbacks
{
    void * (ENET_CALLBACK * malloc) (size_t size);
    void (ENET_CALLBACK * free) (void * memory);
    void (ENET_CALLBACK * no_memory) (void);
} ENetCallbacks;
```

Source: `thirdparty/enet/enet/callbacks.h` — ENetCallbacks

Code examples: ENetCallbacks

Evidence:
- Code: `thirdparty/enet/enet/callbacks.h` — ENetCallbacks
- Code: `thirdparty/enet/callbacks.c` — enet_initialize_with_callbacks
- External (primary, unverified (source anchor not found)): [N1570 — Committee Draft, Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ templates and typed containers

Godot-specific Vector and Ref types are used as C++ template applications throughout this partition.

C++ template applications express typed containers and handles, including Vector collections of Ref-managed connections and parameter-pack-sized argument arrays.

### How it works

#### Typed graph connection collection

GraphEdit returns a Vector whose element type is Ref<GraphEdit::Connection>.

#### Variadic call storage

Node header code sizes an array from sizeof...(p_args), showing a parameter-pack-based helper.

#### Typed scene-state stack

Property utilities pass a Vector<SceneState::PackState> through their comparison and packing logic.

### Anchored code example

```
const Vector<Ref<GraphEdit::Connection>> &GraphEdit::get_connections() const {
```

Source: `scene/gui/graph_edit.cpp` — GraphEdit::get_connections

Code examples: GraphEditConnection, SceneState

Evidence:
- Code: `scene/gui/graph_edit.cpp` — GraphEdit::get_connections
- Code: `scene/main/node.h` — sizeof...(p_args) argument-pointer arrays
- Code: `scene/property_utils.cpp` — Vector<SceneState::PackState> states_stack
- External (primary, verified): [Working Draft, Programming Languages — C++: Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

## Objective-C++ Apple platform adapters

macOS and visionOS implementation files use .mm sources where C++ engine classes meet Apple platform APIs.

Objective-C++ implementation files combine C++ platform adapters with Cocoa-style objects, display coordinates, events, menus, views, and text-to-speech calls.

### How it works

#### DisplayServerMacOS::_get_hdr_output

The macOS display server implements HDR-output lookup in an Objective-C++ source file.

#### KeyMappingMacOS

macOS key mapping accesses keyboard-layout data and maps keycodes, key symbols, and key locations.

#### NativeMenuMacOS

The macOS native-menu adapter stores MenuData and interacts with NSMenu and NSMenuItem objects.

### Anchored code example

```
const NSPoint mouse_pos = [NSEvent mouseLocation];
```

Source: `platform/macos/display_server_macos.mm` — DisplayServerMacOS::mouse_get_position

Prerequisites: cpp-engine-polymorphism

Evidence:
- Code: `platform/macos/display_server_macos.mm` — DisplayServerMacOS::mouse_get_position
- Code: `platform/macos/key_mapping_macos.mm` — KeyMappingMacOS
- Code: `platform/macos/native_menu_macos.h` — NativeMenuMacOS
- External (official, unverified (source anchor not found)): [About Objective-C](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html), accessed 2026-07-15
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

## Python build source generation

The functions produce text and byte-oriented C++ output rather than editor runtime behavior.

Python build helpers perform build-time source generation by transforming inputs into generated C++ source and invoking external translation tooling when available.

### How it works

#### Documentation path table

doc_data_class_path_builder reads and orders paths before emitting a C++ class-path table.

#### Exporter registration

register_exporters_builder emits includes and registration code from platform input.

#### Compressed documentation

make_doc_header compresses documentation data and emits constexpr C++ data.

#### Translation compilation

make_translations uses temporary files and subprocess execution for msgfmt before emitting translation tables.

### Anchored code example

```
def make_translations(target, source, env):
```

Source: `editor/editor_builders.py` — make_translations

Code examples: EditorTranslationList

Evidence:
- Code: `editor/editor_builders.py` — doc_data_class_path_builder
- Code: `editor/editor_builders.py` — register_exporters_builder
- Code: `editor/editor_builders.py` — make_doc_header
- Code: `editor/editor_builders.py` — make_translations
- External (official, unverified (source anchor not found)): [subprocess — Subprocess management — Python 3.14.6 documentation](https://docs.python.org/3/library/subprocess.html), accessed 2026-07-15

## Python SCons module configuration

These scripts are the build-time composition layer for the supplied modules.

Python SCons scripts declare module buildability, dependencies, cloned environments, and source-file collection.

### How it works

#### Module eligibility hook

The SVG config defines can_build and declares jpg and webp module dependencies before returning True.

#### Source collection

OpenXR and scene SCsub scripts create a module object list, add matching C++ sources, and append it to env.modules_sources.

#### Third-party source selection

The Theora and TinyEXR scripts conditionally assemble bundled third-party sources and add them to build environments.

### Anchored code example

```
def can_build(env, platform):
    env.module_add_dependencies("svg", ["jpg", "webp"], True)
    return True
```

Source: `modules/svg/config.py` — can_build

Evidence:
- Code: `modules/svg/config.py` — can_build
- Code: `modules/openxr/extensions/SCsub` — module_obj and env.modules_sources
- Code: `modules/theora/SCsub` — thirdparty_sources
- External (official, verified): [Python 3 Language Reference](https://docs.python.org/3/reference/compound_stmts.html#function-definitions), accessed 2026-07-16

## C++ templates and const access

The code uses Ref<T>, Vector<T>, LocalVector<T>, TypedArray<T>, and const references or pointers throughout module and scene implementations.

C++ templates and const-qualified values provide typed access to engine handles and collection views.

### How it works

#### Typed resource handle

VisualShader language integration stores a typed Ref<VisualShader> handle from its shader input.

#### Read-only varying access

VisualShader exposes varyings through a pointer-to-const lookup API.

#### Read-only navigation path access

NavigationAgent3D returns its current path as a const Vector<Vector3> reference.

### Anchored code example

```
const Ref<VisualShader> shader = p_shader;
```

Source: `modules/visual_shader/editor/visual_shader_language_plugin.cpp` — VisualShaderLanguagePlugin

Code examples: VisualShaderGraph, NavigationAgent3D

Evidence:
- Code: `modules/visual_shader/editor/visual_shader_language_plugin.cpp` — VisualShaderLanguagePlugin
- Code: `modules/visual_shader/visual_shader.h` — VisualShader::get_varying_by_index
- Code: `scene/3d/navigation/navigation_agent_3d.h` — NavigationAgent3D::get_current_navigation_path
- External (primary, verified): [Working Draft, Standard for Programming Language C++](https://eel.is/c++draft/temp), accessed 2026-07-16

## Objective-C++ iOS adapters

The iOS SCons build compiles `.mm` sources for device metrics, display, view, main, and OS behavior.

Objective-C++ iOS adapters implement an embedded platform layer for the iOS embedded adapter.

### How it works

#### platform/ios/SCsub

The iOS library source list includes multiple `.mm` files.

#### DisplayServerIOS

The iOS display-server class is declared as a final subclass of `DisplayServerAppleEmbedded`.

#### display_server_ios.mm

The implementation source contains a `utsname` system-information value.

### Anchored code example

```
struct utsname systemInfo;
```

Source: `platform/ios/display_server_ios.mm` — systemInfo

Evidence:
- Code: `platform/ios/SCsub` — ios_lib
- Code: `platform/ios/display_server_ios.h` — DisplayServerIOS
- Code: `platform/ios/display_server_ios.mm` — systemInfo
- External (official, verified): [About Objective-C](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html), accessed 2026-07-16

## CMake native source indexing

The CMake file explicitly labels itself non-functional and configures it for Android Studio editor support.

CMake source-index configuration exposes the native C/C++ source tree to Android Studio as part of Android Gradle assembly.

### How it works

#### CMAKE_CXX_STANDARD

The native-source-index CMake file sets the requested C++ standard to 14.

#### file(GLOB_RECURSE SOURCES)

The file recursively collects C-family source files below the Godot root for the editor-support target.

#### add_executable

The source-index target is declared with the collected source and header lists.

### Anchored code example

```
set(CMAKE_CXX_STANDARD 14)
```

Source: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — CMAKE_CXX_STANDARD

Evidence:
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — CMAKE_CXX_STANDARD
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — file(GLOB_RECURSE SOURCES)
- Code: `platform/android/java/nativeSrcsConfigs/CMakeLists.txt` — add_executable
- External (official, unverified (source anchor not found)): [cmake-language(7)](https://cmake.org/cmake/help/latest/manual/cmake-language.7.html), accessed 2026-07-15

## C++ templates for binary data operations

The libraries use templates to retain typed interfaces while parsing binary font tables and copying fixed-size byte sequences.

C++ templates specialize reusable byte operations for a compile-time size and typed OpenType offsets.

### How it works

#### unaligned_copy<S>

The LZ4 helper takes the copy size as a non-type template parameter and invokes memcpy for that fixed size.

#### OffsetTo<Type>

HarfBuzz's OpenType layer uses templated offsets to represent references from binary-table locations to typed objects.

### Anchored code example

```
template<int S>
inline
void unaligned_copy(void * d, void const * s) {
  ::memcpy(d, s, S);
}
```

Source: `thirdparty/graphite/src/inc/Compression.h` — unaligned_copy

Code examples: hb_blob_t, hb_face_t

Evidence:
- Code: `thirdparty/graphite/src/inc/Compression.h` — unaligned_copy
- Code: `thirdparty/harfbuzz/src/hb-open-type.hh` — OffsetTo, UnsizedArrayOf, ArrayOf, and CFFIndex
- External (primary, unverified (source anchor not found)): [C++ working draft: Templates](https://eel.is/c++draft/temp), accessed 2026-07-15

## Groovy Gradle build logic

The Android Gradle root, app, library, and editor build files contain executable Groovy configuration.

Groovy Gradle scripts generate Android build tasks and select variants as part of Android Gradle assembly.

### How it works

#### generateBuildTasks

The root Android build script declares a function with flavor, edition, and Android-distribution parameters.

#### getSconsTaskName

The same script declares a helper that names SCons tasks from flavor, build type, and ABI inputs.

#### app/config.gradle

The app configuration script begins by creating a Groovy map.

### Anchored code example

```
def generateBuildTasks(String flavor = "template", String edition = "standard", String androidDistro = "android") {
```

Source: `platform/android/java/build.gradle` — generateBuildTasks

Evidence:
- Code: `platform/android/java/build.gradle` — generateBuildTasks
- Code: `platform/android/java/build.gradle` — getSconsTaskName
- Code: `platform/android/java/app/config.gradle` — map
- External (official, unverified (source anchor not found)): [The Apache Groovy programming language - Syntax](https://groovy-lang.org/syntax.html), accessed 2026-07-15

## C++ templates and traits

Generic core containers and binding helpers are primarily header-defined template code.

The implementation uses C++ templates and trait specializations to adapt behavior to types, including zero construction, hashing, tuple recursion, and argument conversion.

### How it works

#### Zero-construction trait

AABB specializes is_zero_constructible as true, providing a type-level property consumed by generic code.

#### Hash capability detection

Hash functions define has_hash_method and a HashMapHasherDefault implementation selected through template parameters.

#### Recursive tuple representation

Tuple recursively inherits Tuple<Rest...>, with TupleGet specializations for indexed access.

### Anchored code example

```
struct is_zero_constructible<AABB> : std::true_type {};
```

Source: `core/math/aabb.h` — is_zero_constructible<AABB>

Code examples: AABB

Evidence:
- Code: `core/math/aabb.h` — is_zero_constructible<AABB>
- Code: `core/templates/hashfuncs.h` — has_hash_method
- Code: `core/templates/tuple.h` — Tuple
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ template-based method binding

This keeps native method registration type-aware while producing the reflective MethodBind representation used by script-facing APIs.

ClassDB uses C++ templates, `if constexpr`, and member-function traits to generate binders for native methods.

### How it works

#### Typed member binding

ClassDB::bind_method creates a MethodBind from a member-function pointer and detects raw Object pointer return types with if constexpr.

#### Variadic argument metadata

D_METHOD builds debug argument-name metadata from a variadic template parameter pack.

#### Build-gated registration

GDREGISTER_CLASS conditionally instantiates class registration through GD_IS_CLASS_ENABLED.

### Anchored code example

```
#define GDREGISTER_CLASS(m_class) \
	if constexpr (GD_IS_CLASS_ENABLED(m_class)) { \
		::ClassDB::register_class<m_class>(); \
	}
```

Source: `core/object/class_db.h` — GDREGISTER_CLASS

Prerequisites: cpp-native-classes

Code examples: ClassInfo, Variant

Evidence:
- Code: `core/object/class_db.h` — ClassDB::bind_method
- Code: `core/object/class_db.h` — D_METHOD
- Code: `core/object/class_db.h` — GDREGISTER_CLASS
- External (official, verified): [The Standard — Standard C++](https://isocpp.org/std/the-standard), accessed 2026-07-16

## C++ templates and specialization

The implementation uses type and non-type template parameters instead of a single runtime-polymorphic kernel for all widths and primitive forms.

C++ templates parameterize BVH branching, packet width, vector width, primitive layout, and builder behavior at compile time.

### How it works

#### BVH branching factor

BVHN<N> and its builders use N as a compile-time branching factor, including builder settings that assign branchingFactor from N.

#### Width-specific API layouts

RTCRayNt, RTCHitNt, and RTCRayHitNt provide packet layouts parameterized by compile-time N.

#### Specialized builder dispatch

The two-level builder header declares specializations such as MortonBuilder<4,TriangleMesh,Triangle4> and SAHBuilder variants for concrete mesh and primitive combinations.

### Anchored code example

```
settings.branchingFactor = N;
      settings.maxDepth = BVH::maxBuildDepthLeaf;
      return BVHBuilderBinnedSAH::build<NodeRef>
```

Source: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build

Code examples: BVHN, RTCRayNt, MortonBuilder

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh.h` — BVHN
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h` — RTCRayNt
- Code: `thirdparty/embree/kernels/bvh/bvh_builder_twolevel_internal.h` — MortonBuilder
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ class templates and specialization

Embree passes the lane-width parameter through vector types such as Vec3vf<K> and specializes subgrid intersectors for width four.

C++ class templates parameterize Embree vector and intersector code by lane width, and a fixed-width specialization selects a width-specific implementation; C++ classes are needed to explain class templates.

### How it works

#### Lane-parameterized vector arguments

The packet intersector accepts Vec3vf<K> ray vectors, binding the vector width to template parameter K.

#### Fixed-width subgrid path

The SubGridQuadMIntersector1MoellerTrumbore<4,filter> declaration provides a width-four specialization path.

### Anchored code example

```
struct SubGridQuadMIntersector1MoellerTrumbore<4,filter>
```

Source: `thirdparty/embree/kernels/geometry/subgrid_intersector_moeller.h` — SubGridQuadMIntersector1MoellerTrumbore<4,filter>

Prerequisites: cpp-classes-inheritance

Code examples: TriangleMi, SubGrid

Evidence:
- Code: `thirdparty/embree/kernels/geometry/triangle_intersector_moeller.h` — MoellerTrumboreIntersectorK
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector_moeller.h` — SubGridQuadMIntersector1MoellerTrumbore<4,filter>
- External (primary, verified): [C++ Working Draft — Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

## C++ ownership and polymorphic trees

This is the ownership structure behind lazily composed CSG values.

C++ class templates support CSG nodes through shared ownership, enable_shared_from_this, derived leaf and operation classes, and the Manifold implementation boundary.

### How it works

#### CsgNode

CsgNode inherits enable_shared_from_this, allowing a shared node to recover shared ownership of itself.

#### CsgLeafNode and CsgOpNode

Leaf and operation classes derive from CsgNode to represent alternative tree-node roles.

#### Manifold::Impl

The public Manifold API declares an internal Impl type whose detailed mesh state is defined in the implementation header.

### Anchored code example

```
class CsgNode : public std::enable_shared_from_this<CsgNode> {
```

Source: `thirdparty/manifold/src/csg_tree.h` — class CsgNode

Prerequisites: cpp-generic-math

Code examples: Manifold

Evidence:
- Code: `thirdparty/manifold/src/csg_tree.h` — class CsgNode; class CsgLeafNode; class CsgOpNode
- Code: `thirdparty/manifold/include/manifold/manifold.h` — Manifold::Impl declaration
- Code: `thirdparty/manifold/src/impl.h` — struct Manifold::Impl
- External (primary, unverified (source anchor not found)): [C++ Working Draft N4950](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ runtime symbol loading

Metal, MetalFX, and QuartzCore private headers use this pattern to bind symbols at runtime.

C++ wrapper types resolve framework symbols dynamically through `dlsym`-based helper templates.

### How it works

#### MTL::Private::LoadSymbol

The helper casts a `dlsym` result to the requested pointer type.

#### MTLFX::Private::LoadSymbol

The MetalFX private header repeats the dynamic symbol-loading mechanism for its namespace.

### Anchored code example

```
const _Type* pAddress = static_cast<_Type*>(dlsym(RTLD_DEFAULT, pSymbol));
```

Source: `thirdparty/metal-cpp/Metal/MTLPrivate.hpp` — MTL::Private::LoadSymbol

Prerequisites: cpp-native-wrappers

Code examples: MTL4::Archive

Evidence:
- Code: `thirdparty/metal-cpp/Metal/MTLPrivate.hpp` — MTL::Private::LoadSymbol
- Code: `thirdparty/metal-cpp/MetalFX/MTLFXPrivate.hpp` — MTLFX::Private::LoadSymbol
- External (primary, unverified (source anchor not found)): [C++ Working Draft](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ specialized argument marshalling

The specializations select direct, converting, reference, vector, pointer, object, reference-counted, and typed-container paths.

The binding layer specializes PtrToArg through C++ templates to marshal bound method types and return values.

### How it works

#### Scalar conversion specialization

The boolean specialization converts through uint8_t rather than passing a bool storage pointer directly.

#### Geometry argument choices

Method pointer conversion chooses direct or by-reference transport for individual geometry types.

#### Typed container support

PtrToArg has dedicated typed-array and typed-dictionary specializations.

### Anchored code example

```
struct PtrToArg<bool> : Internal::PtrToArgConvert<bool, uint8_t> {};
```

Source: `core/variant/method_ptrcall.h` — PtrToArg<bool>

Prerequisites: cpp-templates-traits

Code examples: Variant, PropertyInfo

Evidence:
- Code: `core/variant/method_ptrcall.h` — PtrToArg<bool>
- Code: `core/variant/method_ptrcall.h` — PtrToArg<Vector3>
- Code: `core/variant/type_info.h` — GetTypeInfo
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ templates and specialization

The reusable interpolation structure is specialized where quaternion behavior needs its own implementation.

The glTF importer uses C++ class templates for interpolation and supplies a quaternion specialization; accessor decoding also instantiates numeric decoding for concrete element types.

### How it works

#### Quaternion interpolation specialization

SceneFormatImporterGLTFInterpolate has a Quaternion specialization in the glTF document implementation.

#### Typed numeric decoding

GLTFAccessor decodes numeric representations through _decode_as_numbers<double>.

### Anchored code example

```
struct SceneFormatImporterGLTFInterpolate<Quaternion> {
```

Source: `modules/gltf/gltf_document.cpp` — SceneFormatImporterGLTFInterpolate<Quaternion>

Code examples: GLTFAnimation, GLTFAccessor

Evidence:
- Code: `modules/gltf/gltf_document.cpp` — SceneFormatImporterGLTFInterpolate<Quaternion>
- Code: `modules/gltf/structures/gltf_accessor.cpp` — GLTFAccessor::_decode_as_numbers<double>
- External (primary, verified): [C++ Working Draft: Templates](https://eel.is/c++draft/temp), accessed 2026-07-16

## C++ SIMD wrapper specialization

The same vector-oriented algorithms are instantiated against widths such as four, eight, and sixteen lanes, with platform-specific definitions behind the wrapper names.

C++ templates and wrapper types represent fixed-width integer, unsigned-integer, float, mask, and long-vector lanes while ISA-specific headers implement their operations.

### How it works

#### ISA-specific vector classes

vint<4> is implemented in an SSE2 header, vint<8> in AVX and AVX2 headers, and vuint<16> in an AVX-512 header.

#### Lane algorithms

The integer wrappers compose shuffle, min or max, and select operations to sort or rearrange lane values.

#### Traversal use

BVH traversal uses vfloat<N>, vint<N>, and vbool<N> to calculate slab intersections and masks for multiple children or rays.

### Anchored code example

```
const vint8 a0 = v;
const vint8 b0 = shuffle<1,0,3,2>(a0);
const vint8 c0 = umin(a0,b0);
const vint8 d0 = umax(a0,b0);
```

Source: `thirdparty/embree/common/simd/vint8_avx2.h` — sort_ascending

Prerequisites: cxx-templates

Code examples: vint4, vint8, vuint16

Evidence:
- Code: `thirdparty/embree/common/simd/vint4_sse2.h` — vint<4>
- Code: `thirdparty/embree/common/simd/vint8_avx2.h` — vint<8>
- Code: `thirdparty/embree/common/simd/vuint16_avx512.h` — vuint<16>
- Code: `thirdparty/embree/kernels/bvh/node_intersector1.h` — BVHNNodeIntersector1
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C preprocessor platform and precision selection

The same source tree supports multiple sample precisions, compilers, and CPU feature sets through macro-controlled compilation.

C conditional compilation selects precision-dependent names, optional compiler intrinsics, and architecture-specific implementations before the codec modules are compiled.

### How it works

#### jpeg_nbits.h: USE_CLZ_INTRINSIC

Compiler and architecture macros select either a count-leading-zero intrinsic path or a lookup-table path for bit counts.

#### jsamplecomp.h: BITS_IN_JSAMPLE

Precision selection maps generic sample types and API/internal names to their 16-bit or alternate precision-specific forms.

#### libpng/intel/intel_init.c: png_init_filter_functions_sse2

When the SSE2 implementation is enabled, the initializer assigns specialized filter functions to the PNG read-filter table according to bytes per pixel.

### Anchored code example

```
#ifdef USE_CLZ_INTRINSIC
#if defined(_MSC_VER) && !defined(__clang__)
#define JPEG_NBITS_NONZERO(x)  (32 - _CountLeadingZeros(x))
#else
#define JPEG_NBITS_NONZERO(x)  (32 - __builtin_clz(x))
#endif
#define JPEG_NBITS(x)          (x ? JPEG_NBITS_NONZERO(x) : 0)
```

Source: `thirdparty/libjpeg-turbo/src/jpeg_nbits.h` — JPEG_NBITS

Code examples: JPEG Upsampler State

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jpeg_nbits.h` — USE_CLZ_INTRINSIC
- Code: `thirdparty/libjpeg-turbo/src/jsamplecomp.h` — BITS_IN_JSAMPLE
- Code: `thirdparty/libpng/intel/intel_init.c` — png_init_filter_functions_sse2
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming Languages: C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ SIMD wrappers and intrinsics

Embree wraps x86 and ARM SIMD operations, while CVTT selects SSE2 support at preprocessing time.

C++ vector wrappers invoke architecture intrinsics to operate on multiple lane values.

### How it works

#### Vector reciprocal operation

Embree's vfloat4 wrapper invokes _mm_rcp_ps for an SSE reciprocal estimate.

#### Platform-specific vector definitions

Embree chooses SIMD vector specializations through varying.h and architecture-specific headers.

#### SSE2 capability selection

CVTT defines CVTT_USE_SSE2 when compiler or target macros indicate SSE2 availability.

### Anchored code example

```
const vfloat4 r = _mm_rcp_ps(a);
```

Source: `thirdparty/embree/common/simd/vfloat4_sse2.h` — embree::rcp

Prerequisites: cpp-abstraction-and-polymorphism

Evidence:
- Code: `thirdparty/embree/common/simd/vfloat4_sse2.h` — embree::rcp
- Code: `thirdparty/embree/common/simd/varying.h` — embree::vtypes
- Code: `thirdparty/cvtt/ConvectionKernels_Config.h` — CVTT_USE_SSE2
- External (official, verified): [ISO/IEC 14882:2020 — Programming languages — C++](https://www.iso.org/standard/79358.html), accessed 2026-07-16
- External (official, unverified (source anchor not found)): [Intel Intrinsics Guide](https://www.intel.com/content/www/us/en/docs/intrinsics-guide/index.html), accessed 2026-07-15

## Python SCons build hooks

The platform build layer uses Python functions called by SCons environments.

Python build scripts define platform detection, option and flag setup, configuration hooks, bundle generation, template assembly, and Emscripten helper functions.

### How it works

#### platform.web.detect.configure

The Web platform exposes a configure hook alongside platform name, build capability, option, flag, and library-emitter functions.

#### create_engine_file

Web Emscripten helpers create engine output and collect JavaScript libraries, pre/post scripts, and externs.

#### generate_bundle

macOS builder code supplies bundle-generation and debug-app construction functions.

### Anchored code example

```
def generate_bundle(target, source, env):
```

Source: `platform/macos/platform_macos_builders.py` — generate_bundle

Evidence:
- Code: `platform/web/detect.py` — configure
- Code: `platform/web/emscripten_helpers.py` — create_engine_file
- Code: `platform/macos/platform_macos_builders.py` — generate_bundle
- External (official, unverified (source anchor not found)): [The Python Language Reference](https://docs.python.org/3/reference/), accessed 2026-07-15

## Preprocessor-selected platform configuration

Jolt centralizes many feature flags in Core.h and applies branches in platform-specific implementation headers.

Preprocessor configuration selects processor architecture, platform APIs, floating-point precision, instruction-set support, diagnostics, and optional subsystems.

### How it works

#### Build configuration string

GetConfigurationString conditionally emits architecture, word size, SIMD, and feature information.

#### Precision-selected real types

Real.h selects float-based or double-based Real, RVec3, and RMat44 aliases using JPH_DOUBLE_PRECISION.

### Anchored code example

```
#ifdef JPH_DOUBLE_PRECISION

// Define real to double
using Real = double;
using Real3 = Double3;
using RVec3 = DVec3;
```

Source: `thirdparty/jolt_physics/Jolt/Math/Real.h` — JPH_DOUBLE_PRECISION branch

Code examples: Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/ConfigurationString.h` — GetConfigurationString
- Code: `thirdparty/jolt_physics/Jolt/Math/Real.h` — Real, RVec3, RMat44 aliases
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## Macro-based RTTI registration

The JPH_IMPLEMENT_SERIALIZABLE and JPH_ADD_ATTRIBUTE macro families are the repository's reflection convention.

Macro-based RTTI registration attaches class metadata and serializable member attributes to types that participate in dynamic lookup and object streaming.

### How it works

#### Member attribute registration

LinearCurve registers its Point coordinates and points array through serializable attributes.

#### Runtime type lookup

Factory finds RTTI by name or hash and inspects registered primitive member types.

### Anchored code example

```
JPH_IMPLEMENT_SERIALIZABLE_NON_VIRTUAL(LinearCurve::Point)
{
	JPH_ADD_ATTRIBUTE(Point, mX)
	JPH_ADD_ATTRIBUTE(Point, mY)
}
```

Source: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — LinearCurve::Point RTTI registration

Prerequisites: cxx-object-model

Code examples: ConstraintSettings, PhysicsMaterialSimple

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — JPH_IMPLEMENT_SERIALIZABLE_NON_VIRTUAL
- Code: `thirdparty/jolt_physics/Jolt/Core/Factory.cpp` — Factory::Find
- Code: `thirdparty/jolt_physics/Jolt/Core/RTTI.h` — RTTI
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## GDScript query-object API use

The documented API uses object construction and member access rather than an inline positional-argument ray query.

GDScript can configure a physics query object, pass it to a physics space, and receive the collision result through chained method calls.

### How it works

#### Ray parameter construction

PhysicsRayQueryParameters2D.create constructs a query from origin and destination vectors.

#### Direct space query

PhysicsDirectSpaceState2D.intersect_ray accepts the parameter object and returns collision data.

#### Shape query configuration

PhysicsShapeQueryParameters2D groups shape, motion, mask, collision-kind, and exclusion settings before a direct-space query.

### Anchored code example

```
var query = PhysicsRayQueryParameters2D.create(global_position, global_position + Vector2(0, 100))
var collision = get_world_2d().direct_space_state.intersect_ray(query)
```

Source: `doc/classes/PhysicsRayQueryParameters2D.xml` — PhysicsRayQueryParameters2D usage example

Code examples: PhysicsShapeQueryParameters2D, PhysicsTestMotionResult2D

Evidence:
- Code: `doc/classes/PhysicsRayQueryParameters2D.xml` — PhysicsRayQueryParameters2D usage example
- Code: `doc/classes/PhysicsDirectSpaceState2D.xml` — PhysicsDirectSpaceState2D.intersect_ray
- Code: `doc/classes/PhysicsShapeQueryParameters2D.xml` — PhysicsShapeQueryParameters2D
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C preprocessor platform selection

SDL's source files and headers select platform code with feature macros and use extern C guards around C-facing declarations.

C preprocessor conditions select platform implementations and preserve a C linkage boundary for C++ compilation units.

### How it works

#### SDL_build_config.h

The build configuration header centralizes feature and platform definitions used by SDL sources.

#### SDL_xinput.c

The source uses an __cplusplus guard to request C linkage when a C++ unit includes the C API declarations.

### Anchored code example

```
#ifdef __cplusplus
extern "C" {
#endif
```

Source: `thirdparty/sdl/core/windows/SDL_xinput.c` — C linkage guard

Evidence:
- Code: `thirdparty/sdl/include/build_config/SDL_build_config.h` — SDL_build_config_h_
- Code: `thirdparty/sdl/core/windows/SDL_xinput.c` — C linkage guard
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:2024 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n3220.pdf), accessed 2026-07-15

## SIMD wrappers and alignment

Public math types hide platform intrinsics behind Vec3, Vec4, UVec4, matrices, and box helpers.

SIMD vector types use aligned storage and compile-time CPU branches to implement vector, matrix, bounding-box, and geometry operations across SSE, NEON, and RISC-V vector paths.

### How it works

#### RISC-V vector gathers

Vec4 implementations scale per-lane offsets and gather four floating-point values through RVV intrinsics.

#### Four-box overlap testing

AABox4 splats a single box into vector lanes and computes axis separation tests against four boxes at once.

### Anchored code example

```
const vfloat32m1_t gathered = __riscv_vluxei32_v_f32m1(inBase, scaled_offsets, 4);
```

Source: `thirdparty/jolt_physics/Jolt/Math/Vec4.inl` — RVV gather implementation

Prerequisites: cxx-preprocessor-configuration

Code examples: Body

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec4.inl` — RVV gather implementation
- Code: `thirdparty/jolt_physics/Jolt/Geometry/AABox4.h` — AABox4VsBox
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec3.h` — alignas(JPH_VECTOR_ALIGNMENT) Vec3
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## Stream-oriented binary serialization

Several types pair SaveBinaryState and RestoreBinaryState methods with their RTTI registrations.

StreamIn and StreamOut implement binary persistence using RTTI hashes for dynamically typed objects and explicit field writes for concrete state.

### How it works

#### Polymorphic type identity

GroupFilter writes its RTTI hash before subtype-specific state is restored through StreamUtils.

#### Field-wise material persistence

PhysicsMaterialSimple writes and reads its debug name and debug color explicitly.

### Anchored code example

```
void GroupFilter::SaveBinaryState(StreamOut &inStream) const
{
	inStream.Write(GetRTTI()->GetHash());
}
```

Source: `thirdparty/jolt_physics/Jolt/Physics/Collision/GroupFilter.cpp` — GroupFilter::SaveBinaryState

Prerequisites: cxx-reflection-macros

Code examples: ConstraintSettings, PhysicsMaterialSimple, Skeleton

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/StreamIn.h` — StreamIn
- Code: `thirdparty/jolt_physics/Jolt/Core/StreamOut.h` — StreamOut
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/GroupFilter.cpp` — GroupFilter::SaveBinaryState and sRestoreFromBinaryState
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## GDScript declarations and typing

GDScript is implemented in a dedicated module that models compiled scripts, member information, analysis, and compilation.

GDScript source uses `extends`, `class_name`, typed variables, functions, and annotations; the engine represents each script as a GDScript Resource.

### How it works

#### Script runtime representation

GDScript derives from Script and records validity, reload state, and member metadata in the native implementation.

#### Typed collection syntax

The class-reference examples use a typed Dictionary declaration with String keys and int values.

#### Parser implementation

The GDScript module contains a dedicated GDScriptParser implementation for source parsing.

### Anchored code example

```
var typed_dict: Dictionary[String, int] = {
```

Source: `doc/classes/Dictionary.xml` — Dictionary typed dictionary code block

Code examples: GDScript, Variant

Evidence:
- Code: `modules/gdscript/gdscript.h` — GDScript
- Code: `modules/gdscript/gdscript_parser.cpp` — GDScriptParser::parse
- Code: `doc/classes/Dictionary.xml` — Dictionary typed dictionary code block
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/latest/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C++ preprocessor configuration

HarfBuzz uses feature macros such as HB_NO_COLOR and include guards throughout this partition.

Preprocessor conditions select optional HarfBuzz subsystems and header inclusion boundaries before C++ compilation.

### How it works

#### _hb_subset_table_color

The color-table dispatcher is compiled only when HB_NO_COLOR is not defined.

#### hb-raster.hh

The raster header uses a conventional include guard.

### Anchored code example

```
#ifndef HB_NO_COLOR
  switch (tag)
```

Source: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color

Code examples: hb_subset_plan_t

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-table-color.cc` — _hb_subset_table_color
- Code: `thirdparty/harfbuzz/src/hb-raster.hh` — HB_RASTER_HH
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/n4910.pdf), accessed 2026-07-15

## GDScript signals and callbacks

This is the idiomatic script-facing bridge to Godot’s signal and input-dispatch systems.

GDScript connects a Signal to a Callable and uses callback functions such as `_input` in Object-derived nodes.

### How it works

#### Signal connection syntax

The ConfirmationDialog reference connects the pressed signal to a callback with Signal.connect syntax.

#### Input callback convention

InputEvent documentation directs event handling to Node._input, which GDScript examples implement as a function callback.

#### Callable transport

Object and editor APIs accept Callable values for deferred or externally invoked behavior.

### Anchored code example

```
get_cancel_button().pressed.connect(_on_canceled)
```

Source: `doc/classes/ConfirmationDialog.xml` — ConfirmationDialog description code block

Prerequisites: gdscript-declarations

Code examples: Node, Variant

Evidence:
- Code: `doc/classes/ConfirmationDialog.xml` — ConfirmationDialog description code block
- Code: `doc/classes/InputEvent.xml` — InputEvent
- Code: `doc/classes/Object.xml` — Object.bind
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/latest/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## Types, inference, and conversions

This is the foundational type-system topic for the fixture corpus.

GDScript fixtures contrast explicit type hints, `:=` inference, `Variant` boundaries, null assignment, and conversion at typed assignments and returns.

### How it works

#### Typed return conversion

Functions declared to return `float` accept integer literals, arguments, and locals in the conversion cases.

#### Weak local inference

A local declared with `=` can later receive values with different types, while a typed destination receives the resulting value.

### Anchored code example

```
func convert_literal_int_to_float() -> float: return 76
func convert_arg_int_to_float(arg: int) -> float: return arg
```

Source: `modules/gdscript/tests/scripts/analyzer/features/return_conversions.gd` — convert_literal_int_to_float()

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/return_conversions.gd` — convert_literal_int_to_float()
- Code: `modules/gdscript/tests/scripts/analyzer/features/local_inference_is_weak.gd` — test()
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C-compatible guarded headers

This is the boundary mechanism for the platform extension's C-facing declarations.

The VI declaration file uses preprocessor guards and conditionally enters `extern "C"` when compiled as C++, so one declaration set can be included by C and C++ translation units.

### How it works

#### Outer inclusion guard

`VULKAN_VI_H_` prevents the header body from being expanded more than once in one translation unit.

#### Conditional C++ linkage

The `__cplusplus` branch surrounds declarations with `extern "C"` only for C++ compilation.

#### Prototype availability switch

`VK_NO_PROTOTYPES` and `VK_ONLY_EXPORTED_PROTOTYPES` control whether the direct function prototype is declared.

### Anchored code example

```
#ifndef VULKAN_VI_H_
#define VULKAN_VI_H_ 1
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VULKAN_VI_H_

Code examples: VkViSurfaceCreateInfoNN

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VULKAN_VI_H_
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — extern "C"
- External (primary, unverified (source anchor not found)): [N1570 Committee Draft — Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## C++ JNI Variant marshalling

The native Android wrapper identifies Java values through JNI calls and produces Godot `Variant` representations.

C++ JNI Variant marshalling builds on RefCounted adapter classes to convert Java primitive wrappers, strings, arrays, objects, and callables into engine Variant values.

### How it works

#### JavaClassWrapper::_jobject_to_variant

The conversion implementation calls Java wrapper methods such as `Integer.intValue()` through JNI and assigns the result to a Variant.

#### jcallable_to_callable

The conversion source identifies a helper that turns a Java callable into an engine callable representation.

#### callable_jni

The callable JNI source reconstructs Variant argument pointers before executing a native callable.

### Anchored code example

```
var = env->CallIntMethod(obj, JavaClassWrapper::singleton->Integer_integerValue);
```

Source: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant

Prerequisites: cpp-runtime-adapters

Code examples: Callable

Evidence:
- Code: `platform/android/java_class_wrapper.cpp` — JavaClassWrapper::_jobject_to_variant
- Code: `platform/android/java_class_wrapper.cpp` — jcallable_to_callable
- Code: `platform/android/variant/callable_jni.cpp` — callable_jni
- External (primary, unverified (source anchor not found)): [Draft C++ Standard: Contents](https://eel.is/c++draft/), accessed 2026-07-15

## C++ single-header implementation selection

The .cpp file is intentionally minimal because VMA bodies are emitted from the included header.

VMA uses preprocessor configuration to place its header implementation in one C++ translation unit while selecting a non-MSVC debug macro before inclusion.

### How it works

#### Implementation emission

VMA_IMPLEMENTATION is defined before vk_mem_alloc.h is included.

#### Build-mode bridge

When DEBUG_ENABLED is set outside MSVC, the translation unit defines _DEBUG.

#### Godot Vulkan dependency

The patched VMA header includes drivers/vulkan/godot_vulkan.h instead of directly including the upstream Vulkan header.

### Anchored code example

```
#define VMA_IMPLEMENTATION
#ifdef DEBUG_ENABLED
#ifndef _MSC_VER
#define _DEBUG
#endif
#endif
#include "vk_mem_alloc.h"
```

Source: `thirdparty/vulkan/vk_mem_alloc.cpp` — translation-unit implementation configuration

Code examples: VmaAllocatorCreateInfo

Evidence:
- Code: `thirdparty/vulkan/vk_mem_alloc.cpp` — VMA_IMPLEMENTATION
- Code: `thirdparty/vulkan/patches/0002-VMA-godot-vulkan.patch` — vk_mem_alloc.h include replacement
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://eel.is/c++draft/), accessed 2026-07-15

## C++ static ABI contracts

The static-assertion header checks wrapper size compatibility with C Vulkan handles and selected type traits.

The binding uses a macro-selected assertion facility to verify wrapper ABI layout and copy or move properties at compile time.

### How it works

#### Wrapper-size contract

The `Instance` wrapper must have the same size as `VkInstance`.

#### Copy and move contracts

The header checks that `Instance` is copy constructible and nothrow move constructible.

#### Macro-selected assertion facility

`VULKAN_HPP_STATIC_ASSERT` defaults to `static_assert` unless callers configure it differently.

### Anchored code example

```
VULKAN_HPP_STATIC_ASSERT( sizeof( VULKAN_HPP_NAMESPACE::Instance ) == sizeof( VkInstance ), "handle and wrapper have different size!" );
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_static_assertions.hpp` — VK_VERSION_1_0 Instance size assertion

Prerequisites: cpp-preprocessor-configuration

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_static_assertions.hpp` — VK_VERSION_1_0 Instance size, copy-construction, and move-construction assertions
- Code: `thirdparty/vulkan/include/vulkan/vulkan_hpp_macros.hpp` — VULKAN_HPP_STATIC_ASSERT
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4861.pdf), accessed 2026-07-15

## C++ tagged storage and casts

Variant separates the active type tag from storage-handling code for each supported runtime type.

The implementation represents runtime Variant values with a C++ type tag and accesses payloads through explicit casts and type-specific accessors.

### How it works

#### Active type tag

Variant initializes its type member to NIL, establishing an explicit runtime type discriminator.

#### Stored transform access

Variant implementation code accesses stored Transform3D data through a type-specific internal storage member.

#### Type-specific accessors

VariantInternalAccessor has specializations for value, object, reference, packed-array, and typed-container forms.

### Anchored code example

```
Type type = NIL;
```

Source: `core/variant/variant.h` — Variant::type

Code examples: Variant, Transform3D

Evidence:
- Code: `core/variant/variant.h` — Variant::type
- Code: `core/variant/variant.cpp` — Variant::operator Transform3D
- Code: `core/variant/variant_internal.h` — VariantInternalAccessor
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ variadic binding

This lets native method and signal calls share typed, compile-time-generated argument marshalling.

The binding code uses variadic templates and parameter packs to move argument lists into Variant pointer arrays.

### How it works

#### Method-definition names

MethodDefinition materializes a parameter pack into a null-terminated argument-name array.

#### Object call arguments

Object templates allocate Variant argument and pointer arrays sized with sizeof...(p_args).

#### Callable argument arrays

Callable templates likewise form Variant pointer arrays before dispatch.

### Anchored code example

```
const char *args[sizeof...(p_args) + 1] = { p_args..., nullptr }; // +1 makes sure zero sized arrays are also supported.
```

Source: `core/object/class_db.h` — MethodDefinition

Prerequisites: cpp-templates-traits

Code examples: Variant, Object

Evidence:
- Code: `core/object/class_db.h` — MethodDefinition
- Code: `core/object/object.h` — Object::call
- Code: `core/variant/callable.h` — Callable::call
- External (primary, unverified (source anchor not found)): [Working Draft, Programming Languages — C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2025/n5008.pdf), accessed 2026-07-15

## C++ lambdas and callback objects

This lets generic builder code operate on caller-selected node and leaf representations without hard-coding a single output format.

C++ lambdas and callable objects bind BVH builders to allocation, node creation, leaf creation, progress, and completion behavior.

### How it works

#### Leaf-creation closure

BVHNBuilderVirtual creates a capturing lambda that forwards a primitive range and allocator to createLeaf.

#### Builder callback parameters

GeneralBVHBuilder::BuilderT accepts CreateAllocFunc, CreateNodeFunc, UpdateNodeFunc, CreateLeafFunc, and progress callback types.

#### Progress monitor closure

ProgressMonitorClosure stores a Closure for build-progress reporting.

### Anchored code example

```
auto createLeafFunc = [&] (const PrimRef* prims, const range<size_t>& set, const Allocator& alloc) -> NodeRef {
        return createLeaf(prims,set,alloc);
      };
```

Source: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build

Code examples: ProgressMonitorClosure, GeneralBVHBuilder, BVHNBuilderVirtual

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh_builder.cpp` — BVHNBuilderVirtual::BVHNBuilderV::build
- Code: `thirdparty/embree/kernels/builders/bvh_builder_sah.h` — GeneralBVHBuilder::BuilderT
- Code: `thirdparty/embree/kernels/common/builder.h` — ProgressMonitorClosure
- External (primary, unverified (source anchor not found)): [Working Draft, Standard for Programming Language C++](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2023/n4950.pdf), accessed 2026-07-15

## Lambdas and standard algorithms

The implementation uses lambdas for ordered curve lookup and for constructing static sampled geometry.

Lambdas provide inline comparison and initialization behavior when standard algorithms or static data construction need local callable logic.

### How it works

#### Curve segment lookup

LinearCurve::GetValue passes a lambda comparator to std::lower_bound to locate the first point at or above an input coordinate.

#### Static sampled geometry

Vec3::sUnitSphere is initialized by a lambda that constructs a StaticArray of sample vectors.

### Anchored code example

```
Points::const_iterator i2 = std::lower_bound(mPoints.begin(), mPoints.end(), inX, [](const Point &inPoint, float inValue) { return inPoint.mX < inValue; });
```

Source: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — LinearCurve::GetValue

Code examples: LinearCurve

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/LinearCurve.cpp` — LinearCurve::GetValue
- Code: `thirdparty/jolt_physics/Jolt/Math/Vec3.cpp` — Vec3::sUnitSphere
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## Godot shader source composition

The repository distinguishes high-level .gdshader source from RenderingDevice-native GLSL and SPIR-V resources.

Godot's shader language writes Shader resource source and can compose reusable ShaderInclude fragments with a preprocessor include directive.

### How it works

#### Shader source resource

Shader is the Resource class for a custom shader program implemented in the Godot shading language.

#### Include composition

ShaderInclude stores source text intended for a Shader preprocessor include.

#### RenderingDevice stage sources

RDShaderSource stores native stage source text for RenderingDevice instead of the high-level Shader resource.

### Anchored code example

```
#include "res://shader_lib.gdshaderinc"
```

Source: `doc/classes/ShaderInclude.xml` — ShaderInclude description

Code examples: ShaderMaterial, RDPipelineSpecializationConstant

Evidence:
- Code: `doc/classes/Shader.xml` — Shader
- Code: `doc/classes/ShaderInclude.xml` — ShaderInclude
- Code: `doc/classes/RDShaderSource.xml` — RDShaderSource.set_stage_source
- External (official, unverified (source anchor not found)): [Shading language](https://docs.godotengine.org/en/stable/tutorials/shaders/shader_reference/shading_language.html), accessed 2026-07-15

## Typed arrays and dictionaries

Collection shape is tested both during analysis and during runtime assignment or argument passing.

Typed arrays and dictionaries apply static typing to element and key/value shapes; the fixture suite exercises declarations, inference, casts, defaults, transfers, and invalid runtime assignments.

### How it works

#### Custom-class array element

`Array[Inner]` stores an inner-class instance and returns it through an `Inner`-typed local.

#### Typed-dictionary failure paths

Runtime error cases distinguish untyped dictionaries from dictionaries whose key and value types are constrained.

### Anchored code example

```
class Inner:
	var prop = "Inner"


var array: Array[Inner] = [Inner.new()]
```

Source: `modules/gdscript/tests/scripts/analyzer/features/typed_array_with_custom_class.gd` — Inner and array

Prerequisites: types-inference-and-conversions

Code examples: Test Script Fixture, Expected Result Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/typed_array_with_custom_class.gd` — array
- Code: `modules/gdscript/tests/scripts/analyzer/features/typed_dictionary_with_custom_class.gd` — dict
- Code: `modules/gdscript/tests/scripts/runtime/errors/typed_dictionary.out` — TypedDictionary.Key and TypedDictionary.Value errors
- External (official, unverified (source anchor not found)): [Array class reference](https://docs.godotengine.org/en/stable/classes/class_array.html), accessed 2026-07-15
- External (official, unverified (source anchor not found)): [Dictionary class reference](https://docs.godotengine.org/en/stable/classes/class_dictionary.html), accessed 2026-07-15

## Warnings and selective suppression

Warning text is itself captured as a tested output contract.

The analyzer reports unsafe, unused, shadowing, inference, conversion, and coroutine cases, while `@warning_ignore` selectively suppresses specified warnings.

### How it works

#### Unsafe argument suppression

A local suppression precedes a call whose argument may have an unsafe static type.

#### Targeted warning fixtures

Separate fixtures capture unsafe casts, unused identifiers, confusable names, and await diagnostics.

### Anchored code example

```
@warning_ignore("unsafe_call_argument")
		print(check(arg))
```

Source: `modules/gdscript/tests/scripts/analyzer/features/null_initializer.gd` — check_arg()

Prerequisites: types-inference-and-conversions

> Prerequisite evidence for `types-inference-and-conversions` needs additional evidence: the declared quote did not appear verbatim in this topic description.

Code examples: Expected Result Fixture, Test Script Fixture

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/null_initializer.gd` — check_arg()
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/unsafe_call_argument.out` — UNSAFE_CALL_ARGUMENT
- Code: `modules/gdscript/tests/scripts/analyzer/warnings/confusable_capture_reassignment.out` — CONFUSABLE_CAPTURE_REASSIGNMENT
- External (official, verified): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-16

## C++ export callbacks

EditorExportPlatform carries callback types so common packing code can delegate destination-specific work.

C++ function-pointer and callable declarations connect generic packaging code to save, removal, shared-object, and script callback implementations.

### How it works

#### Save callback type

EditorExportSaveFunction defines the function-pointer boundary used by platform packaging operations.

#### Script callback data

ScriptCallbackData groups callback data used when export processing invokes script-provided behavior.

#### Packaging entry points

The platform's pack and ZIP operations accept callback-oriented save paths instead of hard-coding every output destination.

### Anchored code example

```
struct ScriptCallbackData {
```

Source: `editor/export/editor_export_platform.h` — EditorExportPlatform::ScriptCallbackData

Code examples: EditorExportPlatform

Evidence:
- Code: `editor/export/editor_export_platform.h` — EditorExportSaveFunction
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::ScriptCallbackData
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::save_pack
- External (primary, verified): [Draft C++ Standard — Function declarators](https://eel.is/c++draft/dcl.fct), accessed 2026-07-16

## C++ const references and pointers

This is a pervasive access pattern in rendering, scene, GUI, and resource code.

C++ const references and pointers expose existing repository objects and buffers without copying them, such as shader parameter names, vectors, arrays, and geometry data.

### How it works

#### Name lookup return pointer

CanvasItem returns a pointer to an existing StringName remapping entry.

#### Connection list borrow

GraphEdit returns its connection collection by const reference.

#### Buffer access

Animation and geometry code obtains typed pointers from packed container storage through ptr().

### Anchored code example

```
const StringName *CanvasItem::_instance_shader_parameter_get_remap(const StringName &p_name) const {
```

Source: `scene/main/canvas_item.cpp` — CanvasItem::_instance_shader_parameter_get_remap

Code examples: Animation, GraphEditConnection

Evidence:
- Code: `scene/main/canvas_item.cpp` — CanvasItem::_instance_shader_parameter_get_remap
- Code: `scene/gui/graph_edit.cpp` — GraphEdit::get_connections
- Code: `scene/resources/animation.cpp` — typed ptr() accesses for animation tracks and compression
- External (primary, verified): [Working Draft, Programming Languages — C++: References](https://eel.is/c++draft/dcl.ref), accessed 2026-07-16

## C++ visitor-style intermediate-tree traversal

glslang uses specialized TIntermTraverser subclasses for reflection, textual output, live-variable gathering, link validation, and limit checking.

C++ inheritance provides traversal specializations that analyze intermediate shader-tree nodes.

### How it works

#### Reflection traversal

TReflectionTraverser is a TIntermTraverser subclass used by reflection.cpp to inspect intermediate declarations and types.

#### Other traversal passes

TOutputTraverser, TVarGatherTraverser, TIOTraverser, and TInductiveTraverser demonstrate separate intermediate-tree analyses and outputs.

### Anchored code example

```
class TReflectionTraverser : public TIntermTraverser {
```

Source: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp` — TReflectionTraverser

Prerequisites: cpp-classes-and-inheritance

Code examples: TObjectReflection

Evidence:
- Code: `thirdparty/glslang/glslang/MachineIndependent/reflection.cpp` — TReflectionTraverser
- Code: `thirdparty/glslang/glslang/MachineIndependent/intermOut.cpp` — TOutputTraverser
- Code: `thirdparty/glslang/glslang/MachineIndependent/iomapper.cpp` — TVarGatherTraverser
- External (primary, unverified (source anchor not found)): [C++ working draft: Classes](https://eel.is/c++draft/class), accessed 2026-07-15

## GDScript typed collections and signature compatibility

The analyzer test corpus specifies accepted and rejected typed containers, function variance, and override signatures.

This extends type declarations with typed Array and Dictionary element constraints, plus parameter and return compatibility checks for inherited functions.

### How it works

#### Typed array argument

The analyzer test declares a function parameter as Array[int].

#### Typed dictionary diagnostics

Tests expect errors when dictionary key or value types conflict with Dictionary[int, int].

#### Override compatibility

Feature tests cover contravariant parameter and covariant return cases.

### Anchored code example

```
func expect_typed(typed: Array[int]):
```

Source: `modules/gdscript/tests/scripts/analyzer/errors/typed_array.gd` — expect_typed

Prerequisites: gdscript-type-declarations

Code examples: GDScriptParser::Node

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/errors/typed_array.gd` — expect_typed
- Code: `modules/gdscript/tests/scripts/analyzer/features/function_param_type_contravariance.gd` — function parameter compatibility cases
- External (official, unverified (source anchor not found)): [Static typing in GDScript](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/static_typing.html), accessed 2026-07-15

## GLSL compute shaders

These shaders are converted into C++ headers by the module build.

Betsy shader sources declare compute workgroup sizes, bind sampled and storage images, fetch texels by invocation ID, and store generated output.

### How it works

#### Workgroup declaration

alpha_stitch.glsl sets an 8 by 8 by 1 local workgroup.

#### Resource bindings

The shader binds two unsigned sampler inputs and one rgba32ui output image.

#### Invocation-indexed writes

main fetches source texels using gl_GlobalInvocationID.xy and stores a combined block.

### Anchored code example

```
void main() {
	uvec2 rgbBlock = texelFetch(srcRGB, ivec2(gl_GlobalInvocationID.xy), 0).xy;
	uvec2 alphaBlock = texelFetch(srcAlpha, ivec2(gl_GlobalInvocationID.xy), 0).xy;

	imageStore(dstTexture, ivec2(gl_GlobalInvocationID.xy), uvec4(rgbBlock.xy, alphaBlock.xy));
}
```

Source: `modules/betsy/alpha_stitch.glsl` — main

Code examples: Image

Evidence:
- Code: `modules/betsy/alpha_stitch.glsl` — main
- Code: `modules/betsy/SCsub` — GLSL_HEADER
- External (official, unverified (source anchor not found)): [The OpenGL Shading Language, Version 4.60.8](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.pdf), accessed 2026-07-15

## C function-pointer API declarations

The generated Vulkan headers provide both PFN typedefs and optionally visible direct prototypes.

C function-pointer declarations expose addressable Vulkan entry points and callback-based frame/event integration without requiring a concrete implementation type.

### How it works

#### Dispatchable creation entry point

PFN_vkCreateWaylandSurfaceKHR specifies the complete call signature for the extension function.

#### Conditional direct declaration

The same header declares vkCreateWaylandSurfaceKHR only when VK_NO_PROTOTYPES and VK_ONLY_EXPORTED_PROTOTYPES do not suppress it.

### Anchored code example

```
typedef VkResult (VKAPI_PTR *PFN_vkCreateWaylandSurfaceKHR)(VkInstance instance, const VkWaylandSurfaceCreateInfoKHR* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — PFN_vkCreateWaylandSurfaceKHR

Code examples: VkWaylandSurfaceCreateInfoKHR, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_wayland.h` — PFN_vkCreateWaylandSurfaceKHR
- Code: `thirdparty/wslay/wslay/wslay.h` — wslay_frame_callbacks and wslay_event_callbacks
- External (primary, unverified (source anchor not found)): [ISO/IEC 9899:201x Committee Draft N1570 — Programming languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## C++ const-reference accessors

Metadata and API accessors make read-only intent explicit in both parameter and return types.

C++ member functions expose lookup results through const pointers and accept immutable String references to avoid copying lookup inputs.

### How it works

#### Profile metadata lookup

The two-argument get_io_path takes both path strings by const reference and returns a const IOPath pointer.

#### Wrapper registry access

get_registered_extension_wrappers returns a const Vector reference, giving callers read-only access to the registered wrapper collection.

### Anchored code example

```
const OpenXRInteractionProfileMetadata::IOPath *OpenXRInteractionProfileMetadata::get_io_path(const String &p_interaction_profile, const String &p_io_path) const {
```

Source: `modules/openxr/action_map/openxr_interaction_profile_metadata.cpp` — OpenXRInteractionProfileMetadata::get_io_path

Code examples: OpenXRInteractionProfile

Evidence:
- Code: `modules/openxr/action_map/openxr_interaction_profile_metadata.cpp` — OpenXRInteractionProfileMetadata::get_io_path
- Code: `modules/openxr/openxr_api.cpp` — OpenXRAPI::get_registered_extension_wrappers
- External (primary, verified): [Working Draft, Programming Languages — C++](https://eel.is/c++draft/dcl.fct), accessed 2026-07-16

## C++ constructor and destructor resource lifetime

The Font constructor establishes borrowed references and per-glyph cache storage; its destructor frees that storage.

A C++ class constructor and destructor control the Font's allocation and release of its cached advance array.

### How it works

#### Font::Font initialization

The constructor stores the face reference, derives scale from units per em, selects font-operation callbacks, allocates advances, and fills them with an invalid sentinel.

#### Font::~Font cleanup

The destructor releases the dynamically allocated advances array.

### Anchored code example

```
Font::Font(float ppm, const Face & f, const void * appFontHandle, const gr_font_ops * ops)
: m_appFontHandle(appFontHandle ? appFontHandle : this),
  m_face(f),
  m_scale(ppm / f.glyphs().unitsPerEm()),
  m_hinted(appFontHandle && ops && (ops->glyph_advance_x || ops->glyph_advance_y))
{
```

Source: `thirdparty/graphite/src/Font.cpp` — Font::Font

Prerequisites: cpp-classes-and-inheritance

Code examples: Font

Evidence:
- Code: `thirdparty/graphite/src/Font.cpp` — Font::Font
- Code: `thirdparty/graphite/src/Font.cpp` — Font::~Font
- External (primary, unverified (source anchor not found)): [C++ working draft: Constructors](https://eel.is/c++draft/class.ctor), accessed 2026-07-15
- External (primary, unverified (source anchor not found)): [C++ working draft: Destructors](https://eel.is/c++draft/class.dtor), accessed 2026-07-15

## C ABI records and dispatch signatures

The record is supplied by a caller; the dispatch signature names instance, creation input, allocation callbacks, and an output surface pointer.

`VkViSurfaceCreateInfoNN` and `PFN_vkCreateViSurfaceNN` express ABI-shaped call parameters through a tagged C record and a typed creation-function pointer.

### How it works

#### Tagged input record

The record contains `sType`, `pNext`, flags, and the platform window pointer.

#### Dispatchable function type

The `PFN_` typedef gives the creation call a reusable typed function-pointer signature.

### Anchored code example

```
typedef VkResult (VKAPI_PTR *PFN_vkCreateViSurfaceNN)(VkInstance instance, const VkViSurfaceCreateInfoNN* pCreateInfo, const VkAllocationCallbacks* pAllocator, VkSurfaceKHR* pSurface);
```

Source: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — PFN_vkCreateViSurfaceNN

Prerequisites: c-compatible-header-guards

Code examples: VkViSurfaceCreateInfoNN, VkSurfaceKHR

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h` — PFN_vkCreateViSurfaceNN
- External (primary, unverified (source anchor not found)): [N1570 Committee Draft — Programming Languages — C](https://www.open-std.org/jtc1/sc22/wg14/www/docs/n1570.pdf), accessed 2026-07-15

## RAII, non-copyable resources, and intrusive references

The library explicitly marks many resources NonCopyable and wraps longer-lived objects in intrusive reference types.

RAII-bound C++ objects manage locks and scopes, while Ref and RefTarget keep shared simulation objects alive without ordinary copying.

### How it works

#### Mutex-scoped recording

DebugRendererRecorder creates a lock_guard before modifying its current frame or serialized command stream.

#### Job lifetime ownership

JobHandle privately derives from Ref<Job>, so a handle retains its Job while dependency and queue operations proceed.

### Anchored code example

```
lock_guard lock(mMutex);
```

Source: `thirdparty/jolt_physics/Jolt/Renderer/DebugRendererRecorder.cpp` — DebugRendererRecorder::DrawLine

Prerequisites: cxx-object-model

Code examples: Shape, PhysicsScene

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Core/Reference.h` — RefTarget, Ref, RefConst
- Code: `thirdparty/jolt_physics/Jolt/Core/JobSystem.h` — JobHandle : private Ref<Job>
- Code: `thirdparty/jolt_physics/Jolt/Renderer/DebugRendererRecorder.cpp` — DebugRendererRecorder::DrawLine
- External (official, verified): [ISO/IEC 14882:2024 — Programming languages — C++](https://www.iso.org/standard/83626.html), accessed 2026-07-16

## GDScript Android plugin integration

The instrumented Godot project is a runnable scene with test scripts for Android plugins, storage, and Java class wrapping.

GDScript test code retrieves JNISingleton objects, calls JavaClassWrapper, and connects signals through Android plugin signals.

### How it works

#### signal_tests.gd::_init

The signal test keeps a `JNISingleton` plugin reference.

#### signal_tests.gd::test_signal_connection

The test connects two named plugin signals to GDScript callback methods.

#### java_class_wrapper_tests.gd::test_interface_object_proxy

The test constructs a GDScript object and passes it to `JavaClassWrapper.create_proxy`.

### Anchored code example

```
var _plugin: JNISingleton
```

Source: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd` — _plugin

Code examples: SignalInfo

Evidence:
- Code: `platform/android/java/app/src/instrumented/assets/test/android_plugin/signal_tests.gd` — test_signal_connection
- Code: `platform/android/java/app/src/instrumented/assets/test/javaclasswrapper/java_class_wrapper_tests.gd` — test_interface_object_proxy
- Code: `platform/android/java/app/src/instrumented/assets/main.gd` — _on_gd_script_toast_button_pressed
- External (official, unverified (source anchor not found)): [GDScript reference](https://docs.godotengine.org/en/stable/tutorials/scripting/gdscript/gdscript_basics.html), accessed 2026-07-15

## Cycle resolutions

- Removed `c-abi-data-structures → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cplusplus-export-callbacks → gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-closures → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-signals-callables → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `csharp-unsafe-interop → c-aggregate-callback-modules`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-signals-callables → csharp-partial-source-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-closures → gdscript-signals-callables`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-lambdas-standard-algorithms → callables-and-lambdas`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-signals-callables → signals-and-await`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-query-objects → c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-query-objects → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-declarations → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `gdscript-declarations → gdscript-query-objects`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-code-generation → python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-code-generation → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-scripts → python-build-actions`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-scripts → python-build-code-generation`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-scripts → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-build-hooks → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-build-hooks → python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-build-hooks → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-build-hooks → python-scons-module-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-scripts → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-module-configuration → python-build-scripts`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-scons-module-configuration → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-code-generation → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-preprocessor-configuration → c-preprocessor-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-preprocessor-configuration → cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-const-borrowing → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `c-preprocessor-platform-selection → cxx-preprocessor-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-engine-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-generic-containers → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-template-binding → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-template-binding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-const-access → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-specialization → cpp-class-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-specialization → cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-views → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-specialization → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-for-binary-data → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-traits → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-templates → cxx-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-preprocessor-configuration → c-platform-selection`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-preprocessor-configuration → c-preprocessor-portability`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-template-binding → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-template-binding → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-template-binding → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-template-binding → cpp-templates-traits`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-templates → cpp-generic-containers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-templates → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-generic-containers → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-and-const-access → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-for-binary-data → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-for-binary-data → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-traits → cpp-templates`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-traits → cpp-templates-and-const-access`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates-traits → cpp-templates-for-binary-data`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `objective-cpp-ios-adapters → objective-cpp-apple-adapters`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-abstraction-and-polymorphism → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-inheritance-interfaces`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cpp-native-wrappers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-inheritance → cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-abstraction-and-polymorphism → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-abstraction-and-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-inheritance-interfaces → cxx-object-model`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-abstraction-and-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-class-hierarchy → classes-inheritance-and-lookup`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-class-hierarchy → cpp-abstraction-and-polymorphism`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-class-hierarchy → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-class-hierarchy → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-class-hierarchy → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-object-hierarchies → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-object-model → cxx-class-hierarchy`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-templates → cpp-templates-and-views`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-registration → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-state-and-composition → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-numeric-value-operations → cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-copy-on-write-containers → cpp-class-state-and-composition`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-copy-on-write-containers → cpp-numeric-value-operations`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-and-ref-handles → cpp-basis-transcoding`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-and-ref-handles → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `classes-inheritance-and-lookup → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `classes-inheritance-and-lookup → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `classes-inheritance-and-lookup → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-inheritance → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-inheritance → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-inheritance → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-runtime-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-interface-polymorphism → cpp-class-registration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-interface-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-interface-polymorphism → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `javascript-browser-ffi → javascript-web-runtime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-classes-and-ref-handles → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-resource-and-polymorphism → cpp-classes-and-ref-handles`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `c-pointers-arrays-and-strides → c-abi-manual-lifetime`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cxx-c-abi → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-actions → python-build-hooks`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-actions → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-class-registration → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `cpp-basis-transcoding → cpp-native-classes`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `c-pointers-arrays-and-strides → c-structures-and-pointers`: removed to preserve the stronger inferred prerequisite hierarchy.
- Removed `python-build-hooks → python-scons-configuration`: removed to preserve the stronger inferred prerequisite hierarchy.
<!-- rope-ladder:end document -->
