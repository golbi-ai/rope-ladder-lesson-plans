<!-- rope-ladder:begin document 53c3513d7f31921701d422764b4fe93c3de784839de4b6966ff1659a89eb9f0a -->
# Entity reference

Only persistent, externally exchanged, or cross-cutting structures appear here.

## AABB

Axis-aligned three-dimensional bounds value represented by a position and size.

Role: spatial bounds value

Lifecycle: Created as a geometry value, transformed or merged for queries, passed through APIs or Variant conversion, and discarded as a value object.

| Field | Type | Meaning |
| --- | --- | --- |
| `position` | `Vector3` | Box origin position. |
| `size` | `Vector3` | Box extents. |

Relationships:
- AABB → **TriangleMesh** (one-to-many): TriangleMesh acceleration nodes partition bounds associated with mesh triangles.

Evidence:
- Code: `core/math/aabb.h:38` — AABB
- Code: `core/math/aabb.cpp:339` — AABB::merge
- Code: `core/math/triangle_mesh.h` — TriangleMesh::BVH

## AABBTreeBuilder

Jolt builder for an AABB tree over indexed triangles.

Role: Builds and exposes tree nodes and source triangles.

Lifecycle: Populated during tree construction, queried for nodes and triangles after construction, then released with the builder.

| Field | Type | Meaning |
| --- | --- | --- |
| `mNodes` | `Array<Node>` | Stores constructed tree nodes. |
| `mTriangles` | `Array<IndexedTriangle>` | Stores indexed source triangles. |
| `mMaxTrianglesPerLeaf` | `uint` | Limits triangles placed in each leaf. |

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.h:35` — AABBTreeBuilder
- Code: `thirdparty/jolt_physics/Jolt/AABBTree/AABBTreeBuilder.cpp:7` — AABBTreeBuilder

## APKExportData

Android-export data structure defined by the Android export platform.

Role: carries export-operation data

Lifecycle: Used during an Android export operation; the supplied index does not expose its member layout or retention rules.

Evidence:
- Code: `platform/android/export/export_plugin.h:78` — APKExportData
- Code: `platform/android/export/export_plugin.h:64` — EditorExportPlatformAndroid

## AccessibilityElement

Cross-cutting accessibility node record managed by the AccessKit accessibility server.

Role: Represents an accessibility element that can be retrieved by resource ID and traversed as a child element.

Lifecycle: Created and owned by the accessibility server's RID owner, retrieved during window accessibility processing, and removed when its owner retires it.

Relationships:
- AccessibilityElement → **AccessibilityElement** (one-to-many): An accessibility element can be resolved as another element's child.

Evidence:
- Code: `drivers/accesskit/accessibility_server_accesskit.h:45` — struct AccessibilityElement
- Code: `drivers/accesskit/accessibility_server_accesskit.cpp` — rid_owner.get_or_null and child_ae lookup

## ApplicationInfo

Application and engine metadata record referenced by Vulkan instance configuration.

Role: supplies optional application metadata to instance setup

Lifecycle: A caller populates the record and an `InstanceCreateInfo` may point to it; no creation or retirement operation for this record is declared in the supplied partition.

| Field | Type | Meaning |
| --- | --- | --- |
| `pApplicationName` | `const char *` | Optional application-name pointer. |
| `pEngineName` | `const char *` | Optional engine-name pointer. |
| `pNext` | `const void *` | Optional extension-chain pointer. |

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:6883` — struct ApplicationInfo

## Array

Dynamic ordered container of Variant elements with optional element-type validation.

Role: runtime sequence container

Lifecycle: Created or copied as a Variant-compatible container, mutated through array operations, and released with its private storage.

| Field | Type | Meaning |
| --- | --- | --- |
| `_p` | `ArrayPrivate *` | Private shared container state. |

Relationships:
- Array → **Variant** (one-to-many): An Array stores an ordered sequence of Variant elements.

Evidence:
- Code: `core/variant/array.h:47` — Array
- Code: `core/variant/array.cpp:45` — ArrayPrivate
- Code: `core/variant/container_type_validate.h:43` — ContainerTypeValidate

## AssemblyHasScriptsAttribute

Assembly-level metadata that identifies whether script types require lookup or are supplied directly.

Role: script type inventory

Lifecycle: Generated or constructed for an assembly, retained in assembly metadata, inspected by ScriptManagerBridge, and discarded with the assembly.

| Field | Type | Meaning |
| --- | --- | --- |
| `RequiresLookup` | `bool` | Indicates that script types must be found by lookup instead of the ScriptTypes list. |
| `ScriptTypes` | `Type[]?` | Optional collection of managed types that implement Godot scripts. |

Relationships:
- AssemblyHasScriptsAttribute → **ScriptPathAttribute** (one-to-many): One assembly-level script inventory can describe many script types, each of which can carry script-path metadata.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/AssemblyHasScriptsAttribute.cs:13` — AssemblyHasScriptsAttribute
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ScriptManagerBridge.cs:20` — ScriptManagerBridge

## AudioStreamInteractive Transition

A transition rule addressed by a pair of clip keys.

Role: Interactive-music edge

Lifecycle: A transition is stored in an AudioStreamInteractive transition map and retrieved during playback selection.

| Field | Type | Meaning |
| --- | --- | --- |
| `TransitionKey` | `nested struct` | Key type formed from the source and destination clip key values. |

Evidence:
- Code: `modules/interactive_music/audio_stream_interactive.h` — Transition and TransitionKey
- Code: `modules/interactive_music/audio_stream_interactive.cpp:202` — TransitionKey(keys[i].x, keys[i].y)

## BVHN

N-ary internal acceleration hierarchy.

Role: Stores node and leaf forms traversed to narrow ray queries to candidate primitives.

Lifecycle: Allocated and built from primitive references, optionally refit or rotated, traversed for queries, and retired with its reference-owned acceleration data.

| Field | Type | Meaning |
| --- | --- | --- |
| `primTy` | `const PrimitiveType*` | Primitive representation stored in the hierarchy. |

Relationships:
- BVHN → **PrimRef** (one-to-many): One hierarchy is built from many primitive-reference records.

Evidence:
- Code: `thirdparty/embree/kernels/bvh/bvh.h:42` — BVHN
- Code: `thirdparty/embree/kernels/bvh/bvh.h:92` — primTy
- Code: `thirdparty/embree/kernels/bvh/bvh_refit.h:13` — BVHNRefitter

## BasisFileHeader

Serialized .basis-file header that identifies the texture representation and locates slices and auxiliary compressed data.

Role: container metadata

Lifecycle: Present at the start of a .basis byte buffer, interpreted by the Basis transcoder while validating or transcoding, then discarded with that input buffer.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_total_slices` | `packed_uint<3>` | Count used to address the serialized slice-description array. |
| `m_tex_format` | `uint8_t` | Encoded Basis texture-format selector. |
| `m_flags` | `packed_uint<2>` | Serialized header flags, including alpha-slice information. |
| `m_slice_desc_file_ofs` | `packed_uint<4>` | Byte offset of the slice-description array. |

Relationships:
- BasisFileHeader → **BasisSliceDescriptor** (one-to-many): One header addresses the array of descriptors for its compressed slices.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h:99` — basis_file_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:12204` — basisu_transcoder::get_file_info

## BasisSliceDescriptor

Serialized descriptor for one compressed Basis texture slice, including image/level identity, dimensions, and byte range.

Role: slice metadata

Lifecycle: Stored in the .basis descriptor array, selected by slice index during transcoding, and retained only as long as the source buffer is available.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_image_index` | `packed_uint<3>` | Source-image index for the slice. |
| `m_level_index` | `uint8_t` | Mipmap-level index for the slice. |
| `m_orig_width` | `packed_uint<2>` | Original pixel width. |
| `m_orig_height` | `packed_uint<2>` | Original pixel height. |
| `m_file_ofs` | `packed_uint<4>` | Byte offset of the compressed slice payload. |
| `m_file_size` | `packed_uint<4>` | Compressed payload length. |

Relationships:
- BasisSliceDescriptor → **BasisFileHeader** (many-to-one): Each descriptor belongs to the header-addressed descriptor array of one Basis file.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_file_headers.h:32` — basis_slice_desc
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.cpp:12428` — basisu_transcoder::transcode_slice

## BrotliDecoderState

Opaque-to-callers decoding state that stores streaming decoder progress and byte-oriented internal pointers.

Role: stream decompression workspace

Lifecycle: Initialized for decoding, advanced as compressed input is consumed and output is retrieved, then released or reinitialized by the Brotli API.

| Field | Type | Meaning |
| --- | --- | --- |
| `chunks` | `const uint8_t*[16]` | Internal byte-pointer array retained by decoder state. |
| `context_lookup` | `const uint8_t*` | Pointer to context-model lookup data. |

Evidence:
- Code: `thirdparty/brotli/dec/state.h:250` — BrotliDecoderStateStruct
- Code: `thirdparty/brotli/dec/decode.c:2882` — BrotliDecoderTakeOutput

## BuildDiagnostic

Structured build issue displayed by the Godot editor.

Role: build result reporting

Lifecycle: Build tooling writes diagnostics to the issues CSV, BuildProblemsView reads them into BuildDiagnostic values, and the view filters or navigates them.

| Field | Type | Meaning |
| --- | --- | --- |
| `Type` | `DiagnosticType` | Severity classification. |
| `File` | `string?` | Source file associated with the diagnostic. |
| `Line` | `int` | Source line number. |
| `Column` | `int` | Source column number. |
| `Code` | `string?` | Build-tool diagnostic identifier. |
| `Message` | `string` | Diagnostic text. |
| `ProjectFile` | `string?` | Project file associated with the diagnostic. |

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildDiagnostic.cs:3` — BuildDiagnostic
- Code: `modules/mono/editor/GodotTools/GodotTools/Build/BuildProblemsView.cs:59` — ReadDiagnosticsFromFile

## CSGBrush

CSG geometry container with face-oriented data used during CSG construction.

Role: Carries the intermediate brush representation for CSG operations.

Lifecycle: Created from CSG shape inputs, merged or processed by CSG operations, then consumed to produce updated shape geometry.

| Field | Type | Meaning |
| --- | --- | --- |
| `faces` | `Vector<CSGBrush::Face>` | Face collection iterated during mesh merge and CSG processing. |

Evidence:
- Code: `modules/csg/csg.h` — CSGBrush and Face
- Code: `modules/csg/csg_shape.cpp:408` — p_mesh_merge->faces

## Callable

Kotlin wrapper for a native callable pointer.

Role: carries a callable handle across the Android JVM/native boundary

Lifecycle: Created through a private constructor with a native pointer; the supplied index does not expose the release path.

| Field | Type | Meaning |
| --- | --- | --- |
| `nativeCallablePointer` | `Long` | Private native callable pointer stored by the wrapper. |

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/variant/Callable.kt:2` — Callable
- Code: `platform/android/variant/callable_jni.cpp:2` — callable_jni

## CameraFeed

Engine camera-feed abstraction implemented by platform camera modules.

Role: Represents a live camera stream exposed through a CameraServer backend.

Lifecycle: Created by a platform camera backend, receives decoded capture data while active, and is released with the backend or feed.

Evidence:
- Code: `modules/camera/camera_android.h:57` — CameraFeedAndroid : public CameraFeed
- Code: `modules/camera/camera_feed_linux.h:42` — CameraFeedLinux : public CameraFeed

## ColladaParseState

Parser-owned aggregation of COLLADA data referenced by the editor importer.

Role: Preserves source-format objects and track indices while a COLLADA scene is converted.

Lifecycle: Created for a COLLADA parse, filled with source data, consumed by ColladaImport during scene construction, then released with the parser object.

| Field | Type | Meaning |
| --- | --- | --- |
| `camera_data_map` | `HashMap<String, CameraData>` | Camera definitions addressable by COLLADA source ID. |
| `mesh_data_map` | `HashMap<String, MeshData>` | Mesh definitions addressable by COLLADA source ID. |
| `animation_tracks` | `Vector<AnimationTrack>` | Parsed animation tracks addressed by importer code. |
| `referenced_tracks` | `HashMap<String, Vector<int>>` | Track-index groups addressed by track name. |

Evidence:
- Code: `editor/import/3d/collada.h` — Collada::State
- Code: `editor/import/3d/editor_import_collada.cpp:48` — ColladaImport

## CollisionShape2D

Scene node that supplies a Shape2D resource to a CollisionObject2D parent.

Role: Bridges editable scene placement to reusable 2D collision geometry.

Lifecycle: Added beneath a collision object, configured with a shape and enabled state, consumed by physics while active, then removed with its node.

| Field | Type | Meaning |
| --- | --- | --- |
| `shape` | `Shape2D` | Collision geometry resource supplied to the parent. |
| `disabled` | `bool` | Whether collision detection is disabled. |
| `debug_color` | `Color` | Editor and debug collision-shape color. |

Relationships:
- CollisionShape2D → **Shape2D** (many-to-one): Each CollisionShape2D has one assigned Shape2D resource.

Evidence:
- Code: `doc/classes/CollisionShape2D.xml:2` — CollisionShape2D

## ConfigFile

Configurable section-and-key value store.

Role: Persists application configuration values independently of the Resource system.

Lifecycle: Created or loaded, read or updated by section and key, saved to a path, then released.

| Field | Type | Meaning |
| --- | --- | --- |
| `section` | `String` | Configuration section name used for lookup. |
| `key` | `String` | Configuration key within a section. |
| `value` | `Variant` | Stored configuration value. |

Relationships:
- ConfigFile → **Variant** (one-to-many): A configuration file stores Variant values for its section and key entries.

Evidence:
- Code: `doc/classes/ConfigFile.xml:2` — ConfigFile

## ConstraintSettings

Serializable configuration base for a runtime constraint.

Role: Carries enabled state, solver priority and step overrides, drawing scale, and user data before a concrete constraint is created.

Lifecycle: Configured and optionally serialized, used by a concrete settings subtype to construct a runtime Constraint, then retained independently of the runtime solver object.

| Field | Type | Meaning |
| --- | --- | --- |
| `mEnabled` | `bool` | Whether the configured constraint is enabled. |
| `mDrawConstraintSize` | `float` | Scale used for constraint debug drawing. |
| `mConstraintPriority` | `uint` | Priority serialized for solver scheduling. |
| `mNumVelocityStepsOverride` | `uint` | Optional serialized override for velocity solver steps. |
| `mNumPositionStepsOverride` | `uint` | Optional serialized override for position solver steps. |
| `mUserData` | `uint64` | Caller-defined data registered as a serializable attribute. |

Relationships:
- ConstraintSettings → **TwoBodyConstraint** (one-to-one): A concrete two-body settings object creates one corresponding runtime two-body constraint instance.

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/Constraint.h:64` — ConstraintSettings
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/Constraint.cpp` — ConstraintSettings RTTI registration and SaveBinaryState

## DAP::Capabilities

Debug Adapter Protocol capability descriptor exchanged during debugger initialization.

Role: Describes supported protocol features to a debug client.

Lifecycle: Produced during protocol initialization and returned to the client as protocol data.

Evidence:
- Code: `editor/debugger/debug_adapter/debug_adapter_types.h:144` — struct Capabilities
- Code: `editor/debugger/debug_adapter/debug_adapter_protocol.h:72` — class DebugAdapterProtocol : public Object
- External (primary, verified): [Debug Adapter Protocol Overview](https://microsoft.github.io/debug-adapter-protocol/overview.html), accessed 2026-07-16

## DAP::Source

Debug Adapter Protocol source descriptor exchanged by the editor debug adapter.

Role: Identifies a source associated with debugger structures.

Lifecycle: Created or fetched while protocol requests are processed and emitted as protocol data.

Evidence:
- Code: `editor/debugger/debug_adapter/debug_adapter_types.h` — struct Source and struct Breakpoint
- Code: `editor/debugger/debug_adapter/debug_adapter_protocol.h:112` — const DAP::Source &fetch_source(const String &p_path)

## DDSFormatInfo

A DDS format descriptor selected by DDS image loading and saving code.

Role: Maps a DDS format choice to format metadata.

Lifecycle: Defined in the DDS format table, selected for a load or save operation, and discarded with that operation.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `const char *` | Descriptor name initialized by the structure. |

Evidence:
- Code: `modules/dds/dds_enums.h:181` — DDSFormatInfo
- Code: `modules/dds/texture_loader_dds.cpp:148` — dds_format_info[p_dds_format]

## Dictionary

Dynamic key-value container with Variant keys and values plus optional type validation.

Role: runtime associative container

Lifecycle: Created or copied as a Variant-compatible container, updated through keyed operations, and released with its private storage.

| Field | Type | Meaning |
| --- | --- | --- |
| `_p` | `DictionaryPrivate *` | Private shared dictionary state. |

Relationships:
- Dictionary → **Variant** (one-to-many): A Dictionary contains Variant keys and Variant values for its entries.

Evidence:
- Code: `core/variant/dictionary.h:46` — Dictionary
- Code: `core/variant/dictionary.cpp:43` — DictionaryPrivate
- Code: `core/variant/container_type_validate.h:43` — ContainerTypeValidate

## ENetCallbacks

Process-level callback record used by ENet allocation paths.

Role: Supplies allocation, release, and out-of-memory behavior to the ENet runtime.

Lifecycle: Initialized to malloc, free, and abort; optionally replaced by enet_initialize_with_callbacks; used by enet_malloc and enet_free for the process lifetime.

| Field | Type | Meaning |
| --- | --- | --- |
| `malloc` | `ENET_CALLBACK function pointer` | Allocates a requested byte count. |
| `free` | `ENET_CALLBACK function pointer` | Releases an allocated memory pointer. |
| `no_memory` | `ENET_CALLBACK function pointer` | Handles allocation failure. |

Evidence:
- Code: `thirdparty/enet/enet/callbacks.h:15` — ENetCallbacks
- Code: `thirdparty/enet/callbacks.c:2` — callbacks
- Code: `thirdparty/enet/callbacks.c:11` — enet_initialize_with_callbacks

## ENetConnection

Reference-counted wrapper around ENet connection behavior.

Role: Owns connection-level service, compression, and peer-access operations.

Lifecycle: Created by caller code, initialized for ENet use, services events and peers, and is released when references end.

Relationships:
- ENetConnection → **ENetPacketPeer** (one-to-many): One connection can expose many peer wrappers through service and get_peers operations.

Evidence:
- Code: `modules/enet/enet_connection.h:43` — ENetConnection : public RefCounted
- Code: `modules/enet/doc_classes/ENetPacketPeer.xml` — ENetConnection.service and ENetConnection.get_peers description

## ENetHost

Transport endpoint that owns peer capacity and communicates with peers.

Role: Creates and manages the local ENet communication endpoint.

Lifecycle: Created through enet_host_create with a maximum peer count, used for communication, and later retired by the matching host-management lifecycle.

| Field | Type | Meaning |
| --- | --- | --- |
| `peerCount` | `size_t` | Maximum number of peers requested when the host is created. |

Relationships:
- ENetHost → **ENetPeer** (one-to-many): One host allocates capacity for and manages multiple peers.

Evidence:
- Code: `thirdparty/enet/host.c:29` — enet_host_create
- Code: `thirdparty/enet/enet/enet.h:41` — _ENetHost

## ENetMultiplayerPeer

MultiplayerPeer implementation backed by ENet.

Role: Adapts ENet transport behavior to the engine multiplayer-peer interface.

Lifecycle: Created and initialized as a client, server, or mesh by callers, services multiplayer traffic, and is released when no longer referenced.

Relationships:
- ENetMultiplayerPeer → **ENetConnection** (one-to-one): One multiplayer peer uses an ENet connection context for its ENet-backed operation.

Evidence:
- Code: `modules/enet/enet_multiplayer_peer.h:39` — ENetMultiplayerPeer : public MultiplayerPeer
- Code: `modules/enet/doc_classes/ENetMultiplayerPeer.xml` — ENetMultiplayerPeer description

## ENetPacketPeer

PacketPeer implementation representing a peer associated with an ENetConnection.

Role: Exposes peer-level packet interaction through the engine packet-peer abstraction.

Lifecycle: Retrieved from a connection during service or peer enumeration and remains usable while its underlying peer association is valid.

Evidence:
- Code: `modules/enet/enet_packet_peer.h:37` — ENetPacketPeer : public PacketPeer
- Code: `modules/enet/doc_classes/ENetPacketPeer.xml` — ENetPacketPeer description

## ENetPeer

Per-remote-peer transport session state.

Role: Tracks connection state under an ENet host.

Lifecycle: Allocated as part of host peer capacity, changes state while connecting or disconnecting, and ceases with its host or peer teardown.

| Field | Type | Meaning |
| --- | --- | --- |
| `state` | `ENetPeerState` | Current peer connection state changed by the protocol state-transition function. |

Relationships:
- ENetPeer → **ENetHost** (many-to-one): Many peers are managed by one host.

Evidence:
- Code: `thirdparty/enet/enet/enet.h:271` — _ENetPeer
- Code: `thirdparty/enet/protocol.c:36` — enet_protocol_change_state

## EXRImage

TinyEXR in-memory image entity.

Role: Image payload type used by TinyEXR's public EXR API.

Lifecycle: The supplied header indexes EXRImage API usage, but its construction and retirement bodies are not included in the supplied excerpt.

Evidence:
- Code: `thirdparty/tinyexr/tinyexr.h:358` — EXRImage

## EditedScene

Per-open-scene editor record containing editable-scene state and editor-only restoration data.

Role: editor session state

Lifecycle: Held in EditorData while a scene is edited; editor state and selection can be restored from configuration, while retirement behavior is not established by the inspected symbols.

| Field | Type | Meaning |
| --- | --- | --- |
| `root` | `Node *` | Root node of the edited scene. |
| `path` | `String` | Scene path associated with the record. |
| `editor_states` | `Dictionary` | Stored editor plugin state. |
| `selection` | `List<Node *>` | Selected scene nodes. |
| `history_id` | `int` | Identifier linking the scene to undo history. |
| `time_opened` | `uint64_t` | Timestamp recorded for the scene session. |

Evidence:
- Code: `editor/editor_data.h` — EditorData::EditedScene
- Code: `editor/editor_data.cpp:385` — EditorData::load_editor_plugin_states_from_config

## EditorExport

Registry state joining export platforms, presets, plugins, and runnable-preset mappings.

Role: export orchestration state

Lifecycle: Owned by editor coordination code while the editor is active; its collections are populated through export registration and preset operations.

| Field | Type | Meaning |
| --- | --- | --- |
| `export_platforms` | `Vector<Ref<EditorExportPlatform>>` | Registered target platform adapters. |
| `export_presets` | `Vector<Ref<EditorExportPreset>>` | Configured export presets. |
| `export_plugins` | `Vector<Ref<EditorExportPlugin>>` | Registered export extensions. |
| `runnable_presets` | `HashMap<Ref<EditorExportPlatform>, Ref<EditorExportPreset>>` | Preset selected as runnable for each platform. |

Relationships:
- EditorExport → **EditorExportPreset** (one-to-many): One export registry owns many export presets.
- EditorExport → **EditorExportPlatform** (one-to-many): One export registry owns many target platform adapters.

Evidence:
- Code: `editor/export/editor_export.h:38` — EditorExport

## EditorExportPlatform

Reference-counted target export adapter with common packaging facilities and collected messages.

Role: target export behavior

Lifecycle: Registered with EditorExport, queried for target behavior, and used to run project export operations.

| Field | Type | Meaning |
| --- | --- | --- |
| `messages` | `Vector<ExportMessage>` | Messages emitted during export processing. |

Relationships:
- EditorExportPlatform → **EditorExport** (many-to-one): Many target platforms are registered in one export registry.
- EditorExportPlatform → **EditorExportPreset** (one-to-many): One target platform can have many presets.
- EditorExportPlatform → **EditorExportPlatform::ExportMessage** (one-to-many): One target platform can collect many export messages.

Evidence:
- Code: `editor/export/editor_export_platform.h:51` — EditorExportPlatform

## EditorExportPlatform::ExportMessage

Structured export diagnostic containing a message category, severity, and text.

Role: export diagnostic

Lifecycle: Created by platform export processing and retained in the platform message collection until that collection is cleared or replaced.

| Field | Type | Meaning |
| --- | --- | --- |
| `msg_type` | `ExportMessageType` | Message severity classification. |
| `category` | `String` | Diagnostic category. |
| `text` | `String` | Diagnostic text. |

Relationships:
- EditorExportPlatform::ExportMessage → **EditorExportPlatform** (many-to-one): Many diagnostics can be retained by one target platform.

Evidence:
- Code: `editor/export/editor_export_platform.h` — EditorExportPlatform::ExportMessage

## EditorExportPreset

Target export configuration containing file-selection policy, output settings, patches, encryption settings, and property values.

Role: export configuration

Lifecycle: Owned by EditorExport, configured for a platform, and used when a target export is run.

| Field | Type | Meaning |
| --- | --- | --- |
| `platform` | `Ref<EditorExportPlatform>` | Target platform associated with the preset. |
| `export_filter` | `ExportFilter` | Policy determining which resources are exported. |
| `include_filter` | `String` | Include filter text. |
| `exclude_filter` | `String` | Exclude filter text. |
| `export_path` | `String` | Configured output path. |
| `selected_files` | `HashSet<String>` | Explicitly selected project files. |
| `customized_files` | `HashMap<String, FileExportMode>` | Per-file export customization modes. |
| `patches` | `Vector<String>` | Configured patch paths. |
| `values` | `HashMap<StringName, Variant>` | Export option values. |

Relationships:
- EditorExportPreset → **EditorExport** (many-to-one): Many presets are owned by one export registry.
- EditorExportPreset → **EditorExportPlatform** (many-to-one): Many presets can be associated with one target platform.

Evidence:
- Code: `editor/export/editor_export_preset.h:38` — EditorExportPreset

## EditorFileSystemDirectory

Indexed project directory with recursive child directories and file records.

Role: project filesystem topology

Lifecycle: Built or refreshed by filesystem scans and retained as part of the editor filesystem tree.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `String` | Directory name. |
| `modified_time` | `uint64_t` | Recorded directory modification time. |
| `parent` | `EditorFileSystemDirectory *` | Containing directory. |
| `subdirs` | `Vector<EditorFileSystemDirectory *>` | Child directory records. |
| `files` | `Vector<FileInfo *>` | File records directly contained by the directory. |

Relationships:
- EditorFileSystemDirectory → **EditorFileSystemDirectory** (one-to-many): One directory record can contain many child directory records.
- EditorFileSystemDirectory → **EditorFileSystemDirectory::FileInfo** (one-to-many): One directory record can contain many indexed file records.

Evidence:
- Code: `editor/file_system/editor_file_system.h:46` — EditorFileSystemDirectory
- Code: `editor/file_system/editor_file_system.cpp:409` — EditorFileSystem::_scan_filesystem

## EditorFileSystemDirectory::FileInfo

Indexed metadata for one project file, including identity, import state, dependencies, and script-class information.

Role: project file metadata

Lifecycle: Created or updated during filesystem scanning and retained by its containing indexed directory.

| Field | Type | Meaning |
| --- | --- | --- |
| `file` | `String` | File name or path component. |
| `type` | `StringName` | Detected resource type. |
| `uid` | `ResourceUID::ID` | Resource identity value. |
| `modified_time` | `uint64_t` | Recorded source-file modification time. |
| `import_valid` | `bool` | Whether import metadata is valid. |
| `deps` | `Vector<String>` | Stored dependency references. |
| `class_info` | `ScriptClassInfo` | Script-class name, base type, icon, and flags. |

Relationships:
- EditorFileSystemDirectory::FileInfo → **EditorFileSystemDirectory** (many-to-one): Many file records belong to one indexed directory.
- EditorFileSystemDirectory::FileInfo → **EditorFileSystemDirectory::FileInfo** (one-to-many): One file record can store many dependency references that are resolved to current paths when possible.

Evidence:
- Code: `editor/file_system/editor_file_system.h:165` — EditorFileSystemDirectory::FileInfo
- Code: `editor/file_system/editor_file_system.cpp:143` — EditorFileSystemDirectory::get_file_deps

## EditorTranslationList

Generated translation table entry pairing a language code with compiled translation bytes.

Role: generated translation data

Lifecycle: Generated by make_translations during the build and emitted into C++ source for compilation.

| Field | Type | Meaning |
| --- | --- | --- |
| `lang` | `const char *` | Language identifier. |
| `data` | `const unsigned char *` | Generated translation byte data. |

Evidence:
- Code: `editor/editor_builders.py:71` — make_translations

## Expected Result Fixture

A persisted `.out` artifact containing the expected success marker, emitted values, warning text, or error text for a test case.

Role: golden expectation

Lifecycle: Authored beside a fixture, compared during testing, and updated only when the intended observable behavior changes.

| Field | Type | Meaning |
| --- | --- | --- |
| `result class` | `text marker` | A leading marker such as `GDTEST_OK`, `GDTEST_PARSER_ERROR`, or `GDTEST_RUNTIME_ERROR`. |
| `expected content` | `text` | Expected printed output or diagnostic messages in order. |

Relationships:
- Expected Result Fixture → **Test Script Fixture** (one-to-one): Each expected result names the outcome of one same-basename script fixture.

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.out:1` — GDTEST_OK
- Code: `modules/gdscript/tests/scripts/runtime/errors/invalid_cast.out:1` — GDTEST_RUNTIME_ERROR

## FBXDocument

GLTFDocument-derived document processor for FBX input.

Role: Transforms FBX scene data into GLTF-oriented import state and scene content.

Lifecycle: Instantiated for an FBX import, receives buffer or file data, populates an FBXState, and is released after the import operation.

Relationships:
- FBXDocument → **FBXState** (one-to-one): One document conversion operation receives one FBXState object as its working import state.

Evidence:
- Code: `modules/fbx/fbx_document.h:40` — FBXDocument : public GLTFDocument
- Code: `modules/fbx/fbx_document.cpp` — p_state FBX conversion accesses

## FBXState

GLTFState-derived state used while importing FBX data.

Role: Retains FBX-import state and import options used by FBX document processing.

Lifecycle: Created for an FBX import, populated during conversion, supplied to scene-generation paths, and released after import consumers finish.

| Field | Type | Meaning |
| --- | --- | --- |
| `allow_geometry_helper_nodes` | `bool` | Records whether auxiliary geometry-helper nodes were used during import. |

Evidence:
- Code: `modules/fbx/fbx_state.h:39` — FBXState : public GLTFState
- Code: `modules/fbx/doc_classes/FBXState.xml:11` — allow_geometry_helper_nodes

## FTC_SNodeRec

Fixed-capacity small-bitmap cache node.

Role: glyph bitmap cache storage

Lifecycle: The small-bitmap cache creates and uses nodes through its cache callbacks; a node contains a bounded group of bitmap records until cache eviction or cleanup.

| Field | Type | Meaning |
| --- | --- | --- |
| `gnode` | `FTC_GNodeRec` | Embedded generic glyph-cache node metadata. |
| `count` | `FT_UInt` | Number of small-bitmap records represented by the node. |
| `sbits` | `FTC_SBitRec[FTC_SBIT_ITEMS_PER_NODE]` | Fixed array of small bitmap records. |

Evidence:
- Code: `thirdparty/freetype/src/cache/ftcsbits.h:31` — FTC_SNodeRec_
- Code: `thirdparty/freetype/src/cache/ftccback.h:45` — ftc_snode_new

## FT_BZip2FileRec

State record backing an FT_Stream that incrementally decompresses Bzip2 font input.

Role: compressed-stream state

Lifecycle: FT_Stream_OpenBzip2 allocates and initializes it; reads may reset it; ft_bzip2_stream_close finalizes bzlib state, frees the record, and clears the stream descriptor.

| Field | Type | Meaning |
| --- | --- | --- |
| `source` | `FT_Stream` | Parent compressed input stream. |
| `stream` | `FT_Stream` | Embedding decompressed stream. |
| `memory` | `FT_Memory` | FreeType allocator used by the wrapper. |
| `bzstream` | `bz_stream` | libbz2 decompression state. |
| `input` | `FT_Byte[FT_BZIP2_BUFFER_SIZE]` | Compressed input buffer. |
| `buffer` | `FT_Byte[FT_BZIP2_BUFFER_SIZE]` | Decompressed output buffer. |
| `pos` | `FT_ULong` | Current output position. |
| `reset` | `FT_Bool` | Signals that decompression must reset before a later read. |

Relationships:
- FT_BZip2FileRec → **FT_Stream** (one-to-one): Each wrapper retains one source stream and one embedding stream.

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:102` — FT_BZip2FileRec_
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:439` — ft_bzip2_stream_close

## FT_Stream

FreeType byte-stream interface used as the source and embedding stream for compressed font input.

Role: stream I/O abstraction

Lifecycle: A caller supplies a source stream; FT_Stream_OpenBzip2 initializes a target stream's fields and callbacks; its close callback releases only the wrapper state.

| Field | Type | Meaning |
| --- | --- | --- |
| `memory` | `FT_Memory` | Allocator used by the embedding stream and its wrapper. |
| `descriptor.pointer` | `FT_Pointer` | Stores the FT_BZip2File wrapper for an adapted stream. |
| `read` | `FT_Stream_IoFunc` | Optional callback used to obtain source bytes. |
| `base` | `FT_Byte*` | In-memory byte base used when no read callback is present. |
| `close` | `FT_Stream_CloseFunc` | Callback assigned to finalize an adapted stream. |

Relationships:
- FT_Stream → **FT_BZip2FileRec** (one-to-one): An adapted FT_Stream stores one Bzip2 wrapper in descriptor.pointer.

Evidence:
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:471` — FT_Stream_OpenBzip2
- Code: `thirdparty/freetype/src/bzip2/ftbzip2.c:249` — ft_bzip2_file_fill_input

## Face

Graphite2 representation of one font face and its table-access context.

Role: font-table provider

Lifecycle: A face is created from application-provided table operations, supplies decoded font data to fonts and segments, and is released after dependent shaping objects.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_appFaceHandle` | `const void *` | Non-null application face handle retained by the Graphite face. |

Relationships:
- Face → **Font** (one-to-many): Many scaled Font objects can reference the same Face.
- Face → **Segment** (one-to-many): Each shaped segment retains the face used to obtain glyph and layout data.

Evidence:
- Code: `thirdparty/graphite/src/inc/Face.h:202` — graphite2::Face
- Code: `thirdparty/graphite/src/Face.cpp:197` — Face::chooseSilf

## FeatureVal

Graphite feature-value vector associated with a feature map.

Role: shaping feature selection

Lifecycle: It is created for a face's feature map, updated through feature-reference APIs, passed to segment shaping, and released by its caller.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_pMap` | `const FeatureMap *` | Feature map defining the value-vector interpretation. |

Relationships:
- FeatureVal → **Face** (many-to-one): The feature map used by a FeatureVal is obtained from the face's Graphite behavior.
- FeatureVal → **Segment** (one-to-many): Feature values are supplied when shaping segments.

Evidence:
- Code: `thirdparty/graphite/src/inc/FeatureVal.h` — graphite2::FeatureVal and gr_feature_val
- Code: `thirdparty/graphite/src/gr_features.cpp:23` — gr_fref_set_feature_value
- Code: `thirdparty/graphite/src/gr_segment.cpp` — segment feature-value handling

## Font

Graphite2 scaled-font object that supplies advances and glyph metrics for a face.

Role: font metrics provider

Lifecycle: Construction binds a face and application font operations, allocates an advance cache, shaping reads metrics through it, and destruction frees the cache.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_appFontHandle` | `const void * const` | Application font handle or the Font object itself when no handle is supplied. |
| `m_face` | `const Face &` | Face supplying glyph data and units per em. |
| `m_scale` | `float` | Pixels-per-em scale derived from the face's units per em. |
| `m_advances` | `float *` | Per-glyph cached advances initialized to INVALID_ADVANCE. |

Relationships:
- Font → **Face** (many-to-one): Each Font references exactly one Face, while a Face can support many Font instances.

Evidence:
- Code: `thirdparty/graphite/src/inc/Font.h:67` — graphite2::Font
- Code: `thirdparty/graphite/src/Font.cpp` — Font::Font and Font::~Font

## GDExtension

Resource representing a native extension library.

Role: Retains loader and reloadability state for a native extension and its engine integration.

Lifecycle: Loaded as a resource, initialized through its loader, registers extension classes, may reload when allowed, then unloads with extension management.

| Field | Type | Meaning |
| --- | --- | --- |
| `loader` | `Ref<GDExtensionLoader>` | Loader used to access the extension library. |
| `reloadable` | `bool` | Whether the extension may be reloaded. |
| `gdextension` | `ObjectGDExtension` | Underlying native extension integration record. |

Evidence:
- Code: `core/extension/gdextension.h:39` — GDExtension
- Code: `core/object/class_db.h` — ClassDB::register_extension_class

## GDExtensionManifest

Persisted extension configuration file.

Role: Names an initialization symbol, compatibility floor, and platform library.

Lifecycle: Read as .gdextension configuration when the compatibility runner launches Godot against the test project.

| Field | Type | Meaning |
| --- | --- | --- |
| `entry_symbol` | `string` | The manifest sets compatibility_test_init. |
| `compatibility_minimum` | `version string` | The manifest sets 4.1. |
| `libraries` | `platform-to-library-path map` | The manifest maps linux.debug.x86_64 to libcompatibility_test.so. |

Evidence:
- Code: `tests/compatibility_test/godot/compatibility_test.gdextension` — configuration and libraries sections
- Code: `tests/compatibility_test/run_compatibility_test.py:126` — compatibility_check

## GDScriptFunction

Compiled GDScript function data executed by the virtual machine.

Role: Holds function execution metadata and validated operation pointers used by runtime calls.

Lifecycle: Created by bytecode generation, attached to a compiled script, executed by the VM as needed, and cleaned up with the owning script.

| Field | Type | Meaning |
| --- | --- | --- |
| `_default_arg_ptr` | `const int *` | Pointer to default-argument metadata. |
| `_global_names_ptr` | `const StringName *` | Pointer to global-name metadata used by execution. |
| `_operator_funcs_ptr` | `const Variant::ValidatedOperatorEvaluator *` | Pointer to validated operator evaluators. |

Evidence:
- Code: `modules/gdscript/gdscript_function.h` — GDScriptFunction execution pointer fields
- Code: `modules/gdscript/gdscript_byte_codegen.cpp:161` — function = memnew(GDScriptFunction)

## GDScriptInstance

Runtime ScriptInstance that binds a GDScript resource to an object.

Role: Provides instance-level property, method, and RPC behavior for a compiled script.

Lifecycle: Created when a script is instantiated on an object, participates in calls and property access, then is removed from the script instance list during teardown.

| Field | Type | Meaning |
| --- | --- | --- |
| `script` | `Ref<GDScript>` | Associated compiled script accessed by instance operations. |

Relationships:
- GDScriptInstance → **GDScriptFunction** (many-to-many): Instances invoke functions defined by their associated scripts.

Evidence:
- Code: `modules/gdscript/gdscript.h:337` — GDScriptInstance : public ScriptInstance
- Code: `modules/gdscript/gdscript.cpp` — GDScriptInstance method dispatch and script.ptr access

## GDScriptTokenizer::Token

Token entity produced by GDScript tokenization.

Role: Carries token kind and literal information into parsing and editor-language features.

Lifecycle: Produced from source text, consumed by parsing and parser extensions, then discarded or serialized through tokenizer-buffer paths.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `Type` | Token kind initialized to EMPTY. |
| `literal` | `Variant` | Literal value read by language-server parser extension code. |

Evidence:
- Code: `modules/gdscript/gdscript_tokenizer.h` — GDScriptTokenizer::Token
- Code: `modules/gdscript/language_server/gdscript_extend_parser.cpp:130` — token.literal

## GDScriptWorkspace

Reference-counted workspace index for parsed GDScript symbols and language-server queries.

Role: Resolves native, script, parameter, local, and position-based symbols.

Lifecycle: Created for the language-server service, populated from parsed documents, queried by protocol handlers, and released with the service.

Relationships:
- GDScriptWorkspace → **LSP::DocumentSymbol** (one-to-many): One workspace resolves and serves many document symbols.

Evidence:
- Code: `modules/gdscript/language_server/gdscript_workspace.h:39` — GDScriptWorkspace : public RefCounted
- Code: `modules/gdscript/language_server/gdscript_workspace.cpp` — get_native_symbol, get_script_symbol, and resolve_symbol

## GDType

Runtime class metadata containing a type name, superclass relation, hierarchy, constants, enums, and signals.

Role: reflection type descriptor

Lifecycle: Created for a registered class, initialized with inherited and self metadata, and retained for runtime type lookup.

| Field | Type | Meaning |
| --- | --- | --- |
| `super_type` | `const GDType *` | Direct runtime superclass metadata. |
| `name` | `StringName` | Runtime class name. |
| `name_hierarchy` | `Vector<StringName>` | Class-name hierarchy. |

Relationships:
- GDType → **MethodInfo** (one-to-many): GDType exposes signal metadata represented by MethodInfo values.

Evidence:
- Code: `core/object/gdtype.h:38` — GDType
- Code: `core/object/gdtype.h` — GDType::get_signal_map
- Code: `core/object/gdtype.cpp:115` — GDType::add_signal

## GLES3::Texture

Cross-cutting GLES3 texture record.

Role: Stores texture state, including its texture type, for GLES3 rendering operations.

Lifecycle: Created and looked up through TextureStorage ownership paths; the indexed excerpt does not establish exact destruction behavior.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `Type` | Texture type initialized to TYPE_2D. |

Evidence:
- Code: `drivers/gles3/storage/texture_storage.h` — struct Texture and Type type = TYPE_2D
- Code: `drivers/gles3/storage/texture_storage.cpp:1022` — texture_owner.get_or_null(p_texture)

## GLTFBufferView

A resource that identifies a byte range within a glTF buffer.

Role: Binary-data slice

Lifecycle: A buffer view is retrieved from GLTFState during decoding and reads its assigned buffer range; the supplied code does not establish its allocation or retirement path.

| Field | Type | Meaning |
| --- | --- | --- |
| `buffer` | `GLTFBufferIndex` | Index of the source packed-byte buffer. |
| `byte_offset` | `int64_t` | Start offset of the view in the source buffer. |
| `byte_length` | `int64_t` | Length used to calculate the end of the view. |

Evidence:
- Code: `modules/gltf/structures/gltf_buffer_view.cpp:39` — GLTFBufferView::load_buffer_view_data
- Code: `modules/gltf/gltf_state.h` — GLTFState::get_buffer_views

## GLTFMaterial

glTF material definition.

Role: Supplies PBR appearance parameters for mesh primitives.

Lifecycle: Persisted in a GLTFAsset's materials array and loaded with the asset.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `string` | The fixture contains Material1 and Material2. |
| `doubleSided` | `boolean` | Both fixture materials set doubleSided to true. |
| `pbrMetallicRoughness` | `JSON object` | The fixture uses baseColorFactor with metallicFactor or roughnessFactor. |

Evidence:
- Code: `tests/data/models/cube.gltf:66` — materials

## GLTFNode

An indexed scene-node resource with hierarchy and optional scene attachments.

Role: Scene graph record

Lifecycle: Nodes are held in GLTFState and used by GLTFDocument and SkinTool while reconstructing or exporting a scene hierarchy.

| Field | Type | Meaning |
| --- | --- | --- |
| `parent` | `GLTFNodeIndex` | Index of the parent node. |
| `children` | `node-index collection` | Indexes of child nodes. |
| `mesh` | `GLTFMeshIndex` | Optional mesh attachment index. |
| `skin` | `GLTFSkinIndex` | Optional skin attachment index. |
| `skeleton` | `GLTFSkeletonIndex` | Optional skeleton association index. |

Relationships:
- GLTFNode → **GLTFNode** (one-to-many): One node can identify many child nodes.
- GLTFNode → **GLTFSkin** (many-to-one): Many skinned nodes can reference one skin.

Evidence:
- Code: `modules/gltf/structures/gltf_node.h:37` — GLTFNode
- Code: `modules/gltf/gltf_document.cpp` — node parent, children, mesh, skin, and skeleton handling

## GLTFSkin

A resource that associates joint node indexes and inverse bind data with a constructed skeleton.

Role: Skinning definition

Lifecycle: The document parser fills skin information, SkinTool uses it to establish skeleton and bone data, and the resulting Godot skin references are compared during processing.

| Field | Type | Meaning |
| --- | --- | --- |
| `joints` | `node-index collection` | Joint node indexes from the glTF skin definition. |
| `joints_original` | `node-index collection` | Original joint node indexes used while assigning bones. |
| `inverse_binds` | `transform collection` | Inverse bind matrices when supplied by the glTF skin. |
| `skeleton` | `GLTFSkeletonIndex` | Associated skeleton index. |

Relationships:
- GLTFSkin → **GLTFNode** (one-to-many): One skin references many joint nodes.

Evidence:
- Code: `modules/gltf/structures/gltf_skin.h:41` — GLTFSkin
- Code: `modules/gltf/gltf_document.cpp` — skin joints and inverseBindMatrices parsing
- Code: `modules/gltf/skin_tool.cpp` — GLTFSkin joints_original and inverse_binds processing

## GodotIdeMetadata

File-backed endpoint metadata used by IDE clients to find the editor messaging server.

Role: IDE connection discovery

Lifecycle: The editor writes metadata for a messaging endpoint; clients read and monitor it before opening a TCP connection; it is replaced as the endpoint changes.

| Field | Type | Meaning |
| --- | --- | --- |
| `Port` | `int` | TCP port of the messaging endpoint. |
| `EditorExecutablePath` | `string` | Path to the editor executable. |

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/GodotIdeMetadata.cs:5` — GodotIdeMetadata
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Client.cs:167` — ReadMetadataFile

## GodotService.EngineStatus

Engine-status type used by the remote Android engine service protocol.

Role: communicates remote engine state

Lifecycle: Produced or updated by `GodotService` and transported under the `KEY_ENGINE_STATUS` message key; individual status members are not exposed by the supplied index.

Evidence:
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt` — GodotService.EngineStatus
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/GodotService.kt:79` — KEY_ENGINE_STATUS
- Code: `platform/android/java/lib/src/main/java/org/godotengine/godot/service/RemoteGodotFragment.kt:54` — GodotService.EngineStatus

## GodotSoftBody3D

A native 3D collision object that models deformable geometry.

Role: Soft-body simulation object

Lifecycle: The 3D server owns soft bodies by handle; active-soft-body lists feed them to the simulation step, and mesh mappings update their physical representation.

| Field | Type | Meaning |
| --- | --- | --- |
| `Node` | `nested struct` | Physical node record used by faces and links. |
| `Link` | `nested struct` | Connection between physical nodes. |
| `Face` | `nested struct` | Triangle-like surface element that refers to nodes. |

Evidence:
- Code: `modules/godot_physics_3d/godot_soft_body_3d.h:46` — GodotSoftBody3D : public GodotCollisionObject3D
- Code: `modules/godot_physics_3d/godot_space_3d.h` — GodotSpace3D::get_active_soft_body_list

## GraphicsPipelineCreateInfo

Externally exchanged C++ configuration record for graphics-pipeline creation.

Role: assembles shader and fixed-function pipeline state

Lifecycle: A caller populates pointers to state records for a pipeline-creation API call; the supplied declarations contain no pipeline lifetime operation.

| Field | Type | Meaning |
| --- | --- | --- |
| `pStages` | `const PipelineShaderStageCreateInfo *` | Pointer to shader-stage configuration records. |
| `pVertexInputState` | `const PipelineVertexInputStateCreateInfo *` | Pointer to vertex-input state. |
| `pRasterizationState` | `const PipelineRasterizationStateCreateInfo *` | Pointer to rasterization state. |
| `pColorBlendState` | `const PipelineColorBlendStateCreateInfo *` | Pointer to color-blend state. |
| `pNext` | `const void *` | Optional extension-chain pointer. |

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:56895` — struct GraphicsPipelineCreateInfo

## Halfedge

Directed edge record used by the Manifold internal triangle topology.

Role: topological adjacency

Lifecycle: Created while preparing internal mesh topology, updated by mesh operations, and retained by Manifold::Impl until that implementation is discarded.

| Field | Type | Meaning |
| --- | --- | --- |
| `startVert` | `int` | Index of the directed edge's start vertex. |
| `endVert` | `int` | Index of the directed edge's end vertex. |
| `pairedHalfedge` | `int` | Index of the oppositely directed adjacent halfedge. |
| `propVert` | `int` | Index of the associated property vertex. |

Relationships:
- Halfedge → **Halfedge** (one-to-one): Each directed halfedge stores the index of its paired adjacent halfedge.

Evidence:
- Code: `thirdparty/manifold/src/shared.h:119` — struct Halfedge
- Code: `thirdparty/manifold/src/impl.cpp` — PrepHalfedges; pairedHalfedge construction
- Code: `thirdparty/manifold/src/edge_op.cpp` — Halfedge topology edits

## InputEvent

Base resource for externally received or programmatically supplied input events.

Role: Carries input-device and subtype-specific event data into input handling.

Lifecycle: Created by platform input, by API code, or by the engine; dispatched through Input; then released when no longer referenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `device` | `int` | Input device identifier. |
| `window_id` | `int` | Receiving window identifier on InputEventFromWindow subtypes. |
| `position` | `Vector2` | Pointer or gesture position on applicable subtypes. |
| `pressed` | `bool` | Pressed state on applicable button and action subtypes. |

Evidence:
- Code: `doc/classes/InputEvent.xml:2` — InputEvent
- Code: `doc/classes/InputEventFromWindow.xml` — InputEventFromWindow.window_id
- Code: `doc/classes/InputEventMouseButton.xml` — InputEventMouseButton.pressed

## InstanceCreateInfo

Externally exchanged Vulkan instance configuration record.

Role: groups application metadata with enabled layers and extensions

Lifecycle: A caller populates the record before an instance-creation call; no instance-creation or retirement function body appears in the supplied partition.

| Field | Type | Meaning |
| --- | --- | --- |
| `pApplicationInfo` | `const ApplicationInfo *` | Optional application metadata record. |
| `ppEnabledLayerNames` | `const char * const *` | Pointer to enabled layer names. |
| `ppEnabledExtensionNames` | `const char * const *` | Pointer to enabled extension names. |
| `pNext` | `const void *` | Optional extension-chain pointer. |

Relationships:
- InstanceCreateInfo → **ApplicationInfo** (zero-to-one): An instance configuration may reference one application-metadata record through `pApplicationInfo`.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:67360` — struct InstanceCreateInfo
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:6883` — struct ApplicationInfo

## JPEG Decompression Context

Public decompression session state that owns the selected JPEG input and processing modules.

Role: Coordinates a JPEG decode or coefficient-transcode operation.

Lifecycle: An application initializes the context, configures its source and selected modules, reads scanlines or coefficients, then finishes decompression and releases associated memory.

| Field | Type | Meaning |
| --- | --- | --- |
| `src` | `struct jpeg_source_mgr *` | Input-source manager. |
| `master` | `struct jpeg_decomp_master *` | Master decompression controller. |
| `entropy` | `struct jpeg_entropy_decoder *` | Entropy-decoder module. |
| `idct` | `struct jpeg_inverse_dct *` | Inverse-DCT module. |
| `upsample` | `struct jpeg_upsampler *` | Upsampling module. |
| `cconvert` | `struct jpeg_color_deconverter *` | Color-deconversion module. |

Relationships:
- JPEG Decompression Context → **JPEG Upsampler State** (one-to-one): A decompression context points to one selected upsampler module for an active decode.

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jpeglib.h:328` — jpeg_decompress_struct

## JPEG Upsampler State

Private state for separate JPEG upsampling and color-buffer emission.

Role: Buffers color rows and dispatches component-specific resampling methods.

Lifecycle: Initialized for a decompression context, consumes input row groups while tracking output rows, and is discarded with the JPEG context.

| Field | Type | Meaning |
| --- | --- | --- |
| `color_buf` | `_JSAMPARRAY[MAX_COMPONENTS]` | Per-component upsampled row-group buffers or direct aliases to full-size input data. |
| `methods` | `upsample1_ptr[MAX_COMPONENTS]` | Per-component upsampling routines. |
| `next_row_out` | `int` | Next row to emit from the color buffer. |
| `rows_to_go` | `JDIMENSION` | Remaining image rows. |
| `h_expand` | `UINT8[MAX_COMPONENTS]` | Cached horizontal expansion factors for integer expansion. |

Relationships:
- JPEG Upsampler State → **JPEG Decompression Context** (one-to-one): The selected JPEG decompression context owns and invokes this upsampler state.

Evidence:
- Code: `thirdparty/libjpeg-turbo/src/jdsample.h:51` — my_upsampler

## JSON-RPC document

Dictionary-shaped external request, notification, success, or error message handled by JSONRPC.

Role: Represents a method name, parameters, and correlation identifier for JSON-RPC exchange.

Lifecycle: Constructed as a Dictionary, converted to or from JSON by caller code, exchanged over a chosen transport, then discarded or retained by the caller.

| Field | Type | Meaning |
| --- | --- | --- |
| `method` | `String` | Method name for request or notification documents. |
| `params` | `Variant` | Array or dictionary method parameters. |
| `id` | `Variant` | Response-correlation identifier when a response is expected. |

Relationships:
- JSON-RPC document → **Variant** (one-to-many): A JSON-RPC document carries Variant-based parameter and identifier values.

Evidence:
- Code: `doc/classes/JSONRPC.xml` — JSONRPC::make_notification
- Code: `doc/classes/JSONRPC.xml:2` — JSONRPC

## JavaScriptObjectImpl

Web-side engine proxy for a JavaScript object.

Role: Converts properties, method calls, callbacks, and values across the C++ and JavaScript boundary.

Lifecycle: Created through JavaScriptBridge object or interface operations, used through proxy calls and conversions, then released with its engine reference.

| Field | Type | Meaning |
| --- | --- | --- |
| `_js_id` | `int` | Identifier used by JavaScript wrapper calls to locate the proxied object. |

Evidence:
- Code: `platform/web/javascript_bridge_singleton.cpp:70` — JavaScriptObjectImpl

## JoltSpace3D

Jolt-backed runtime physics space with JPH interfaces, active-object tracking, query access, and debug contacts.

Role: Jolt simulation space

Lifecycle: JoltPhysicsServer3D retrieves spaces from its owner storage; JoltSpace3D updates its JPH PhysicsSystem and exposes query and debug-contact data.

| Field | Type | Meaning |
| --- | --- | --- |
| `physics_system` | `JPH::PhysicsSystem` | System updated during a physics step. |
| `debug_contacts` | `PackedVector3Array` | Debug-contact points returned through a const getter. |

Evidence:
- Code: `modules/jolt_physics/spaces/jolt_space_3d.cpp` — JoltSpace3D physics_system update and get_debug_contacts
- Code: `modules/jolt_physics/spaces/jolt_space_3d.h:61` — JoltSpace3D

## KTX2 Private Texture State

Internal KTX2 data needed to locate levels and retain supercompression global data.

Role: Tracks KTX2-specific byte layout after texture construction.

Lifecycle: Allocated with a KTX2 texture, populated while reading or creating its level index and supercompression data, and freed when the texture is destroyed.

| Field | Type | Meaning |
| --- | --- | --- |
| `_supercompressionGlobalData` | `ktx_uint8_t *` | Supercompression global data. |
| `_requiredLevelAlignment` | `ktx_uint32_t` | Required alignment for stored texture levels. |
| `_sgdByteLength` | `ktx_uint64_t` | Length of supercompression global data. |
| `_firstLevelFileOffset` | `ktx_uint64_t` | First image-level file offset for stream-created textures whose image data is not yet loaded. |
| `_levelIndex` | `ktxLevelIndexEntry[1]` | Growable level index whose offsets are relative to image-data start. |

Evidence:
- Code: `thirdparty/libktx/lib/texture2.h:33` — ktxTexture2_private

## KTX2 Texture

A KTX texture object with virtual method tables, protected shared state, and KTX2-specific implementation data.

Role: Represents a loaded or constructed GPU texture container.

Lifecycle: Constructed from create information, bytes, a file, or a stream; accessed through texture methods; optionally transcoded; then destructed through its type-specific implementation.

| Field | Type | Meaning |
| --- | --- | --- |
| `vtbl` | `struct ktxTexture_vtbl *` | Public virtual-method table. |
| `vvtbl` | `struct ktxTexture_vvtbl *` | Internal virtual-method table. |
| `_protected` | `struct ktxTexture_protected *` | Shared protected texture state. |
| `orientation` | `struct ktxOrientation` | Texture orientation metadata. |

Relationships:
- KTX2 Texture → **KTX2 Private Texture State** (one-to-one): A KTX2 texture has one private KTX2 state block that records supercompression and level-index data.

Evidence:
- Code: `thirdparty/libktx/include/ktx.h:280` — ktxTexture
- Code: `thirdparty/libktx/lib/texture2.c:46` — ktxTexture2_vtbl

## KinematicCollision2D

Reference-counted report returned for detected PhysicsBody2D movement collisions.

Role: Provides collision counterpart, position, normal, motion, and remainder data for custom responses.

Lifecycle: Produced when movement detects a collision, inspected by caller code, then released when no references remain.

| Field | Type | Meaning |
| --- | --- | --- |
| `collider` | `Object` | Object involved in the collision. |
| `position` | `Vector2` | Collision position. |
| `normal` | `Vector2` | Collision surface normal. |
| `remainder` | `Vector2` | Motion remaining after the collision. |

Evidence:
- Code: `doc/classes/KinematicCollision2D.xml:2` — KinematicCollision2D

## Ktx2Header

Packed KTX2 container header describing dimensions, layers, faces, levels, supercompression, and metadata regions.

Role: container metadata

Lifecycle: Parsed by ktx2_transcoder::init from caller-provided bytes and retained in the transcoder until clear resets the object.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_identifier` | `uint8_t[12]` | KTX2 file identifier bytes. |
| `m_pixel_width` | `basisu::packed_uint<4>` | Texture width in texels. |
| `m_pixel_height` | `basisu::packed_uint<4>` | Texture height in texels. |
| `m_level_count` | `basisu::packed_uint<4>` | Number of indexed texture levels. |
| `m_supercompression_scheme` | `basisu::packed_uint<4>` | Stored supercompression scheme selector. |
| `m_sgd_byte_offset` | `basisu::packed_uint<8>` | Offset of supercompression global data. |

Relationships:
- Ktx2Header → **Ktx2LevelIndex** (one-to-many): One header declares the count for the per-level index array.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:721` — ktx2_header
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_transcoder::init

## Ktx2LevelIndex

Packed per-level KTX2 byte-range record.

Role: level metadata

Lifecycle: Parsed into the ktx2_transcoder level-index vector during initialization and cleared with its transcoder.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_byte_offset` | `basisu::packed_uint<8>` | Start offset of the level payload. |
| `m_byte_length` | `basisu::packed_uint<8>` | Stored payload length. |
| `m_uncompressed_byte_length` | `basisu::packed_uint<8>` | Declared uncompressed level length. |

Relationships:
- Ktx2LevelIndex → **Ktx2Header** (many-to-one): Each record indexes one level declared by a KTX2 header.

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:741` — ktx2_level_index
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h` — ktx2_transcoder::get_level_index

## Ktx2TranscoderState

Thread-specific KTX2 transcoding state containing Basis transcoder state and a cached uncompressed level buffer.

Role: transcoding workspace

Lifecycle: Created by the caller for transcoding work, reset by clear, and reused across KTX2 level operations.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_transcoder_state` | `basist::basisu_transcoder_state` | Embedded state for ETC1S or supercompressed UASTC transcoding. |
| `m_level_uncomp_data` | `basisu::uint8_vec` | Cached uncompressed level bytes. |
| `m_uncomp_data_level_index` | `int` | Index identifying the cached uncompressed level. |

Evidence:
- Code: `thirdparty/basis_universal/transcoder/basisu_transcoder.h:931` — ktx2_transcoder_state

## LSP::DocumentSymbol

Language-server symbol entity used for classes, members, locals, and navigation results.

Role: Represents a named symbol with a URI and hierarchical children.

Lifecycle: Built from parser information, stored or queried by workspace services, returned in language-server requests, and replaced when the document is reparsed.

| Field | Type | Meaning |
| --- | --- | --- |
| `uri` | `String` | Document URI accessed by parser and workspace code. |
| `children` | `symbol collection` | Nested symbols traversed by parser and workspace code. |

Relationships:
- LSP::DocumentSymbol → **LSP::DocumentSymbol** (one-to-many): One parent symbol can contain many child symbols.

Evidence:
- Code: `modules/gdscript/language_server/godot_lsp.h:218` — DocumentSymbol
- Code: `modules/gdscript/language_server/gdscript_extend_parser.cpp` — DocumentSymbol children traversal

## LSP::TextDocumentItem

Language-server document payload type managed by the GDScript protocol.

Role: Represents a text document exchanged with language-server document handling.

Lifecycle: Received or updated by protocol handling, kept in managed document state, parsed for language features, and removed when the document closes.

Relationships:
- LSP::TextDocumentItem → **LSP::DocumentSymbol** (one-to-many): One parsed text document can provide many symbols through its symbol tree.

Evidence:
- Code: `modules/gdscript/language_server/godot_lsp.h:660` — TextDocumentItem
- Code: `modules/gdscript/language_server/gdscript_language_protocol.cpp:393` — managed_files

## LSR

Language-script-region value used by ICU locale matching.

Role: Provides the compact locale components used with likely-subtag and locale-distance data.

Lifecycle: Constructed from locale components during matching and retained only by the matching structures that own it.

| Field | Type | Meaning |
| --- | --- | --- |
| `language` | `const char *` | Language component. |
| `script` | `const char *` | Script component. |
| `region` | `const char *` | Region component. |

Relationships:
- LSR → **LocaleMatcher** (one-to-many): LocaleMatcher stores pointers to supported LSR values.

Evidence:
- Code: `thirdparty/icu4c/common/lsr.h:17` — LSR
- Code: `thirdparty/icu4c/common/unicode/localematcher.h` — LocaleMatcher::supportedLSRs

## LauncherIcon

Android export structure for a launcher-icon export entry.

Role: carries launcher-icon export metadata

Lifecycle: Created for export processing; the supplied index only exposes its `export_path` member and does not establish later ownership.

| Field | Type | Meaning |
| --- | --- | --- |
| `export_path` | `const char *` | Path field of the launcher-icon export entry. |

Evidence:
- Code: `platform/android/export/export_plugin.h:57` — LauncherIcon

## Locale

ICU locale identifier object with language, script, region, variant, and keyword-oriented accessors.

Role: Represents locale identity passed through locale and internationalization services.

Lifecycle: Constructed from locale data or strings, copied or iterated by locale services, and retired by normal object lifetime.

Relationships:
- Locale → **LocaleMatcher** (one-to-many): A locale matcher evaluates one or more desired and supported Locale objects.
- Locale → **LSR** (one-to-one): Locale matching derives an LSR representation for a locale.

Evidence:
- Code: `thirdparty/icu4c/common/unicode/locid.h:51` — Locale
- Code: `thirdparty/icu4c/common/localematcher.cpp:608` — LocaleMatcher::getBestMatch

## LocaleMatcher

ICU locale-selection service.

Role: Finds the best supported locale for desired locale input.

Lifecycle: Built with supported locales and matching data, queried for matches, then released by normal object lifetime.

| Field | Type | Meaning |
| --- | --- | --- |
| `supportedLocales` | `const Locale **` | References candidate supported locales. |
| `supportedLSRs` | `const LSR **` | References derived language-script-region values. |
| `defaultLocale` | `const Locale *` | References the configured default locale. |

Relationships:
- LocaleMatcher → **Locale** (one-to-many): One matcher stores multiple supported locale candidates.
- LocaleMatcher → **LSR** (one-to-many): One matcher stores multiple supported LSR candidates.

Evidence:
- Code: `thirdparty/icu4c/common/unicode/localematcher.h:28` — LocaleMatcher

## MainFrameTime

Frame-timing output structure used by the main timer synchronization subsystem.

Role: carries per-frame timing results

Lifecycle: Constructed or populated during timer advancement and replaced on a later frame; the supplied index does not expose its member layout or destruction behavior.

Evidence:
- Code: `main/main_timer_sync.h:35` — MainFrameTime
- Code: `main/main_timer_sync.cpp:539` — MainTimerSync::advance

## ManagedCallbacks

Sequential ABI table of native-callable managed function pointers.

Role: native-to-managed callback contract

Lifecycle: Generated callback methods populate the table, native code receives the table during initialization, and it remains valid for the initialized managed runtime.

| Field | Type | Meaning |
| --- | --- | --- |
| `SignalAwaiter_SignalCallback` | `delegate* unmanaged<IntPtr, godot_variant**, int, godot_bool*, void>` | Completes a managed signal awaiter from native signal data. |
| `DelegateUtils_InvokeWithVariantArgs` | `delegate* unmanaged<IntPtr, void*, godot_variant**, int, godot_variant*, void>` | Invokes a managed delegate using native variant arguments. |
| `ScriptManagerBridge_CreateManagedForGodotObjectBinding` | `delegate* unmanaged<godot_string_name*, IntPtr, IntPtr>` | Creates managed state for a native Godot object binding. |

Relationships:
- ManagedCallbacks → **Variant** (one-to-many): One callback table defines many callbacks that receive or return native Variant values.

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Bridge/ManagedCallbacks.cs:8` — ManagedCallbacks
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/NativeInterop/NativeFuncs.cs` — NativeFuncs.Initialize

## Material

Resource base for visual surface properties.

Role: Supplies coloring and shading configuration for geometry.

Lifecycle: Created, imported, or loaded; assigned to one or more mesh surfaces or visual instances; used during rendering; then released.

| Field | Type | Meaning |
| --- | --- | --- |
| `render_priority` | `int` | Relative rendering priority. |
| `next_pass` | `Material` | Optional material used in a following rendering pass. |

Evidence:
- Code: `doc/classes/Material.xml:2` — Material

## MeshGLP

Public templated mesh carrier for indexed triangles, vertex properties, and merge/run metadata.

Role: geometry interchange payload

Lifecycle: Created by a caller, supplied to Manifold construction, and read while the internal mesh representation is built.

| Field | Type | Meaning |
| --- | --- | --- |
| `triVerts` | `std::vector<I>` | Triangle vertex-index sequence. |
| `vertProperties` | `std::vector<Precision>` | Vertex-property sequence. |
| `numProp` | `int` | Number of properties per property vertex. |
| `mergeFromVert` | `std::vector<I>` | Source vertices participating in requested merges. |
| `mergeToVert` | `std::vector<I>` | Target vertices for requested merges. |
| `runTransform` | `std::vector<Precision>` | Per-run transform data. |

Relationships:
- MeshGLP → **Halfedge** (one-to-many): Internal construction converts a mesh's triangle connectivity into many directed halfedges.

Evidence:
- Code: `thirdparty/manifold/include/manifold/manifold.h:113` — struct MeshGLP
- Code: `thirdparty/manifold/src/constructors.cpp` — Manifold constructors accepting MeshGLP
- Code: `thirdparty/manifold/src/impl.h` — MeshGLP field access in Manifold::Impl construction

## Message

Framed IDE-messaging envelope containing request or response identity and content.

Role: protocol envelope

Lifecycle: Created by a peer or client, encoded onto the connection, decoded by MessageDecoder, dispatched, and answered or released.

| Field | Type | Meaning |
| --- | --- | --- |
| `Kind` | `MessageKind` | Whether the envelope is a request or response. |
| `Id` | `string` | Operation identifier used for request dispatch and response matching. |
| `Content` | `MessageContent` | Status and textual body for the message. |

Relationships:
- Message → **MessageContent** (one-to-one): Every Message contains exactly one MessageContent value.

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs:3` — Message
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/MessageDecoder.cs:6` — MessageDecoder

## MessageContent

Status-and-body payload for an IDE message.

Role: protocol payload

Lifecycle: Constructed with a status and body, carried by a Message through transport and dispatch, then consumed by a response awaiter or handler.

| Field | Type | Meaning |
| --- | --- | --- |
| `Status` | `MessageStatus` | Protocol result status. |
| `Body` | `string` | Serialized or textual request and response body. |

Evidence:
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/Message.cs:7` — MessageContent
- Code: `modules/mono/editor/GodotTools/GodotTools.IdeMessaging/ResponseAwaiter.cs` — ResponseAwaiter<T>.SetResult

## MethodInfo

Reflection description of a method or signal signature.

Role: callable signature metadata

Lifecycle: Built during binding or signal registration, queried by reflection and invocation code, and owned by metadata structures.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `StringName` | Method or signal name. |
| `return_val` | `PropertyInfo` | Declared return metadata. |
| `arguments` | `Vector<PropertyInfo>` | Declared argument metadata. |
| `default_arguments` | `Vector<Variant>` | Default values for optional arguments. |

Relationships:
- MethodInfo → **PropertyInfo** (one-to-many): A MethodInfo carries return and argument PropertyInfo descriptions.

Evidence:
- Code: `core/object/method_info.h:48` — MethodInfo
- Code: `core/object/method_info.cpp` — MethodInfo::operator==
- Code: `core/object/class_db.cpp` — ClassDB::bind_method

## MotionProperties

Mutable rigid-body motion state and mass-property configuration.

Role: Stores degrees of freedom, inverse mass and inertia, velocities, and simulation statistics for a Body.

Lifecycle: Associated with a movable body, configured from mass properties and allowed degrees of freedom, updated by simulation, and consulted by collision and constraint code.

| Field | Type | Meaning |
| --- | --- | --- |
| `mAllowedDOFs` | `EAllowedDOFs` | The permitted translation and rotation axes. |
| `mInvMass` | `float` | Inverse mass, set to zero when translation is disallowed. |
| `mInvInertiaDiagonal` | `Vec3` | Diagonal inverse inertia, set to zero when rotation is disallowed. |
| `mLinearVelocity` | `Vec3` | Current linear velocity, including kinematic movement targets. |
| `mAngularVelocity` | `Vec3` | Current angular velocity, including kinematic movement targets. |

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionProperties.cpp:12` — MotionProperties::SetMassProperties
- Code: `thirdparty/jolt_physics/Jolt/Physics/Body/MotionProperties.inl` — MotionProperties::MoveKinematic, ClampLinearVelocity, ClampAngularVelocity

## NavigationPathQueryParameters2D

Configurable 2D navigation-server path request.

Role: Carries query map, positions, layers, and region filters into a navigation path query.

Lifecycle: Created, configured, submitted to NavigationServer2D, optionally reused after result consumption, then released.

| Field | Type | Meaning |
| --- | --- | --- |
| `map` | `RID` | Navigation map queried. |
| `start_position` | `Vector2` | Global path origin. |
| `target_position` | `Vector2` | Global path destination. |
| `excluded_regions` | `RID[]` | Region identifiers excluded from the query. |

Evidence:
- Code: `doc/classes/NavigationPathQueryParameters2D.xml:2` — NavigationPathQueryParameters2D

## NodePath

Path value whose private data separates node-name segments from subname segments and records absolute-path status.

Role: node addressing value

Lifecycle: Parsed or constructed into Data, sliced or compared through path APIs, and released with its shared data.

| Field | Type | Meaning |
| --- | --- | --- |
| `data` | `Data *` | Private path storage. |
| `path` | `Vector<StringName>` | Node-name segments in Data. |
| `subpath` | `Vector<StringName>` | Subname segments in Data. |

Relationships:
- NodePath → **StringName** (one-to-many): A NodePath stores multiple interned name segments.

Evidence:
- Code: `core/string/node_path.h:39` — NodePath
- Code: `core/string/node_path.h` — NodePath::Data
- Code: `core/string/node_path.cpp:264` — NodePath::slice

## Object

Base engine object with identity, dynamic properties, calls, signals, metadata, scripts, and runtime type access.

Role: engine runtime object

Lifecycle: Initialized into ObjectDB, accessed by ObjectID and dynamic APIs while alive, then passes pre-delete and cleanup handling.

| Field | Type | Meaning |
| --- | --- | --- |
| `_instance_id` | `ObjectID` | Engine instance identity. |
| `signal_map` | `HashMap<StringName, SignalData>` | Registered signal definitions and connected slots. |
| `script_instance` | `ScriptInstance *` | Attached script-side implementation, when present. |

Relationships:
- Object → **ObjectID** (one-to-one): Each Object exposes one instance identifier.
- Object → **GDType** (one-to-one): Each Object obtains one runtime GDType view for its class.

Evidence:
- Code: `core/object/object.h:132` — Object
- Code: `core/object/object.h` — Object::SignalData
- Code: `core/object/object.cpp:2103` — Object::get_gdtype

## ObjectID

Value object used to identify engine Object instances.

Role: object identity handle

Lifecycle: Issued for an Object instance, used for lookup while the instance is registered, and becomes unusable after the object retires.

| Field | Type | Meaning |
| --- | --- | --- |
| `id` | `uint64_t` | Numeric identity value carried by the handle. |

Relationships:
- ObjectID → **Object** (one-to-one): An ObjectID identifies one Object instance while that instance is registered.

Evidence:
- Code: `core/object/object_id.h:41` — ObjectID
- Code: `core/object/object.h:874` — ObjectDB

## Ogg Packet

An externally exchanged encoded payload plus stream-position flags and counters.

Role: Transfers codec data between Ogg stream framing and codec APIs.

Lifecycle: Produced from a codec or page stream, consumed by a codec or stream operation, and retained only while its referenced packet bytes remain valid.

| Field | Type | Meaning |
| --- | --- | --- |
| `packet` | `unsigned char *` | Encoded packet bytes. |
| `bytes` | `long` | Packet byte count. |
| `b_o_s` | `long` | Beginning-of-stream flag. |
| `e_o_s` | `long` | End-of-stream flag. |
| `granulepos` | `ogg_int64_t` | Codec-defined time or position marker. |
| `packetno` | `ogg_int64_t` | Packet sequence counter. |

Evidence:
- Code: `thirdparty/libogg/ogg/ogg.h:86` — ogg_packet

## Ogg Stream State

Mutable state for packetizing and depacketizing one logical Ogg stream.

Role: Retains body data, lacing information, page sequencing, and granule progress.

Lifecycle: Initialized for a logical stream, accepts or emits packets and pages over time, then is cleared.

| Field | Type | Meaning |
| --- | --- | --- |
| `body_data` | `unsigned char *` | Accumulated page-body storage. |
| `lacing_vals` | `int *` | Lacing values describing packet segmentation. |
| `serialno` | `long` | Logical-stream serial number. |
| `pageno` | `long` | Current page sequence number. |
| `granulepos` | `ogg_int64_t` | Current codec-defined granule position. |

Relationships:
- Ogg Stream State → **Ogg Packet** (one-to-many): One logical stream state accepts or produces a sequence of packets.

Evidence:
- Code: `thirdparty/libogg/ogg/ogg.h:49` — ogg_stream_state

## OggPacketSequence

Persisted resource holding Ogg packet-sequence payload and timing metadata.

Role: audio packet storage

Lifecycle: Constructed or loaded as a Resource, passed to playback support, and released when resource references end.

| Field | Type | Meaning |
| --- | --- | --- |
| `granule_positions` | `PackedInt64Array` | Granule positions for pages in the packet sequence. |
| `packet_data` | `Array[]` | Raw packets comprising the sequence. |
| `sampling_rate` | `float` | Sample-rate information for the sequence. |

Evidence:
- Code: `modules/ogg/ogg_packet_sequence.h:41` — OggPacketSequence
- External (official, unverified (source anchor not found)): [OggPacketSequence — Godot Engine](https://docs.godotengine.org/en/latest/classes/class_oggpacketsequence.html), accessed 2026-07-15

## OpenXRActionMap

Persisted root resource for the OpenXR input action configuration graph.

Role: XR input configuration

Lifecycle: Authored or loaded as a Resource, consumed by OpenXR action-map code to organize action-related resources, and released with the resource graph.

Evidence:
- Code: `modules/openxr/action_map/openxr_action_map.h:39` — OpenXRActionMap
- Code: `modules/openxr/action_map/openxr_action_map.cpp:37` — OpenXRActionMap

## OpenXRFutureResult

RefCounted object tracking one asynchronous OpenXR future result.

Role: Exposes future identity and result value to callers of asynchronous extension operations.

Lifecycle: An extension creates or registers a future, returns this result object, and callers observe or cancel it until completion.

| Field | Type | Meaning |
| --- | --- | --- |
| `future` | `int` | The related XrFutureEXT value exposed by get_future. |
| `result_value` | `Variant` | Extension-provided asynchronous result exposed by get_result_value. |

Evidence:
- Code: `modules/openxr/extensions/openxr_future_extension.h:57` — class OpenXRFutureResult : public RefCounted
- Code: `modules/openxr/doc_classes/OpenXRFutureResult.xml` — OpenXRFutureResult::get_future and get_result_value

## OpenXRSpatialQueryResultData

Component-data object holding the main results of a spatial snapshot query.

Role: Carries retrieved spatial entity identifiers and tracking states.

Lifecycle: It is supplied as the first requested component for a snapshot query, populated by the query, then read by index.

| Field | Type | Meaning |
| --- | --- | --- |
| `entity_id` | `int` | OpenXR spatial entity identifier returned for an indexed result. |
| `entity_state` | `int` | Tracking state returned for an indexed result. |

Evidence:
- Code: `modules/openxr/extensions/spatial_entities/openxr_spatial_entities.h:221` — class OpenXRSpatialQueryResultData : public OpenXRSpatialComponentData
- Code: `modules/openxr/doc_classes/OpenXRSpatialQueryResultData.xml` — OpenXRSpatialQueryResultData::get_entity_id and get_entity_state

## OpenXRStructureBase

RefCounted wrapper for OpenXR structure data passed into OpenXR APIs.

Role: Represents a structure header and supports next-chain composition.

Lifecycle: A structure object is created, optionally linked to another structure, passed to an API call, and retained while its native data must remain valid.

| Field | Type | Meaning |
| --- | --- | --- |
| `next` | `OpenXRStructureBase` | Head of the next-chain used to extend an API call. |

Relationships:
- OpenXRStructureBase → **OpenXRStructureBase** (one-to-one): Each structure can point to one next structure, forming a linked chain.

Evidence:
- Code: `modules/openxr/openxr_structure.h:39` — class OpenXRStructureBase : public RefCounted
- Code: `modules/openxr/doc_classes/OpenXRStructureBase.xml` — OpenXRStructureBase.next

## PFR_FaceRec

PFR-specific face record extending the generic FreeType face with parsed PFR data.

Role: parsed font face

Lifecycle: Face initialization loads PFR state into the record; the PFR driver subsequently uses the embedded root face and parsed structures for glyph and character-map work.

| Field | Type | Meaning |
| --- | --- | --- |
| `root` | `FT_FaceRec` | Embedded generic FreeType face. |
| `header` | `PFR_HeaderRec` | Parsed PFR file header. |
| `log_font` | `PFR_LogFontRec` | Parsed logical-font data. |
| `phy_font` | `PFR_PhyFontRec` | Parsed physical-font data. |

Evidence:
- Code: `thirdparty/freetype/src/pfr/pfrobjs.h:27` — PFR_FaceRec_
- Code: `thirdparty/freetype/src/pfr/pfrobjs.h:60` — pfr_face_init

## PNG Information State

PNG metadata state shared by reading and writing routines.

Role: Stores image dimensions, validity bits, and computed row layout.

Lifecycle: Created with PNG processing state, populated from file chunks or application setters, queried during decoding or encoding, and released with its PNG structures.

| Field | Type | Meaning |
| --- | --- | --- |
| `width` | `png_uint_32` | Image width. |
| `height` | `png_uint_32` | Image height. |
| `valid` | `png_uint_32` | Bitset recording available PNG metadata. |
| `rowbytes` | `png_size_t` | Computed byte count for an image row. |

Evidence:
- Code: `thirdparty/libpng/pnginfo.h:31` — png_info_def
- Code: `thirdparty/libpng/pngget.c:19` — png_get_valid
- Code: `thirdparty/libpng/pngget.c:40` — png_get_rowbytes

## PackedData::PackedFile

Indexed metadata for one file exposed from a resource pack.

Role: Locates and describes a virtual file for pack-source access.

Lifecycle: Created when a pack source adds a path, retained in PackedData indexes and optional delta-patch lists, then removed by path removal or clearing packed data.

| Field | Type | Meaning |
| --- | --- | --- |
| `pack` | `String` | Containing pack path. |
| `offset` | `uint64_t` | File offset in the pack; zero marks an erased file. |
| `size` | `uint64_t` | Stored file size. |
| `md5` | `uint8_t[16]` | Stored MD5 byte array. |
| `src` | `PackSource *` | Pack source that opens the file. |
| `encrypted` | `bool` | Whether the file is encrypted. |
| `delta` | `bool` | Whether the file represents delta data. |

Evidence:
- Code: `core/io/file_access_pack.h:164` — PackedData::PackedFile

## PackedScene

Resource wrapper for serialized scene state.

Role: Carries scene content as a reusable Resource.

Lifecycle: Created or loaded as a Resource; holds scene state for later instantiation; released through Resource ownership when no longer referenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `state` | `Ref<SceneState>` | Reference-counted backing scene state. |

Relationships:
- PackedScene → **SceneState** (one-to-one): A PackedScene holds one backing SceneState reference.

Evidence:
- Code: `scene/resources/packed_scene.h:248` — class PackedScene : public Resource
- Code: `scene/resources/packed_scene.cpp` — PackedScene scene-state access

## PhysicsMaterialSimple

Serializable concrete physics material carrying debug presentation data.

Role: Provides a named and colored material implementation for shapes.

Lifecycle: Constructed as a PhysicsMaterial subtype, serialized or restored with debug fields, referenced by shapes, and released through intrusive reference ownership.

| Field | Type | Meaning |
| --- | --- | --- |
| `mDebugName` | `String` | Debug-facing material name. |
| `mDebugColor` | `Color` | Debug-facing material color. |

Relationships:
- PhysicsMaterialSimple → **Shape** (one-to-many): One material can be returned by many shapes or subshapes.

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/PhysicsMaterialSimple.cpp` — PhysicsMaterialSimple RTTI registration, SaveBinaryState, RestoreBinaryState
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/PhysicsMaterialSimple.h:12` — PhysicsMaterialSimple

## PhysicsShapeQueryParameters2D

Configures a 2D shape query issued to PhysicsDirectSpaceState2D.

Role: Physics query request.

Lifecycle: Created, populated with query filters and motion, submitted to a direct space-state method, then released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `motion` | `Vector2` | Requested sweep direction and distance. |
| `collision_mask` | `int` | Bitmask of physics layers considered by the query. |
| `exclude` | `RID[]` | Low-level objects omitted from the query. |
| `collide_with_bodies` | `bool` | Whether physics bodies are query candidates. |

Relationships:
- PhysicsShapeQueryParameters2D → **RID** (many-to-many): The exclusion list holds RID values for low-level objects.

Evidence:
- Code: `doc/classes/PhysicsShapeQueryParameters2D.xml:2` — PhysicsShapeQueryParameters2D
- Code: `doc/classes/PhysicsDirectSpaceState2D.xml` — PhysicsDirectSpaceState2D.cast_motion

## PhysicsTestMotionResult2D

Carries the motion and collision outcome of a 2D body motion test.

Role: Physics query result.

Lifecycle: Supplied to or populated by a body_test_motion call, read for collider and motion details, then released when unreferenced.

Relationships:
- PhysicsTestMotionResult2D → **RID** (many-to-one): A collision result can expose the colliding body's RID.

Evidence:
- Code: `doc/classes/PhysicsTestMotionResult2D.xml:2` — PhysicsTestMotionResult2D

## PluginConfigAndroid

Android-plugin configuration structure in the Android exporter partition.

Role: carries Android plugin configuration

Lifecycle: Its construction, serialization, and retirement behavior are not exposed by the supplied index.

Evidence:
- Code: `platform/android/export/godot_plugin_config.h:53` — PluginConfigAndroid
- Code: `platform/android/export/godot_plugin_config.cpp:41` — PluginConfigAndroid

## PolyPathD

Floating-point polygon-tree node with a polygon, scale factor, parent relationship, and owned child nodes.

Role: hierarchical clipping result

Lifecycle: Created as a root or through AddChild, populated during clipping, and clears or destroys its owned child list on reset or destruction.

| Field | Type | Meaning |
| --- | --- | --- |
| `childs_` | `PolyPathDList` | Owned child polygon-tree nodes. |
| `scale_` | `double` | Coordinate conversion scale inherited from the parent when present. |
| `polygon_` | `PathD` | Polygon stored at this tree node. |

Relationships:
- PolyPathD → **PolyPathD** (one-to-many): A node owns zero or more child polygon-tree nodes.

Evidence:
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h:70` — PolyPathD
- Code: `thirdparty/clipper2/include/clipper2/clipper.engine.h` — PolyPathD::AddChild

## PortMapping

A parsed port-mapping result node.

Role: Represents one mapping in a parsed IGD port-mapping listing.

Lifecycle: Port-list parsing allocates and links mapping nodes into parser data; the parser cleanup path retires the resulting list.

| Field | Type | Meaning |
| --- | --- | --- |
| `l_next` | `struct PortMapping *` | Links this mapping to the next parsed mapping. |

Relationships:
- PortMapping → **PortMapping** (one-to-many): A parser result can contain many mappings joined by `l_next`.

Evidence:
- Code: `thirdparty/miniupnpc/include/miniupnpc/portlistingparse.h:48` — struct PortMapping
- Code: `thirdparty/miniupnpc/src/portlistingparse.c:149` — ParsePortListing

## PresentInfoKHR

Externally exchanged presentation request record.

Role: submits swapchain image presentation with synchronization inputs

Lifecycle: A caller supplies the record to a presentation operation; the record is not itself a persistent resource.

| Field | Type | Meaning |
| --- | --- | --- |
| `pWaitSemaphores` | `const Semaphore *` | Pointer to synchronization objects waited before presentation. |
| `pSwapchains` | `const SwapchainKHR *` | Pointer to target swapchain handles. |
| `pImageIndices` | `const uint32_t *` | Pointer to image indices corresponding to target swapchains. |
| `pNext` | `const void *` | Optional extension-chain pointer. |

Relationships:
- PresentInfoKHR → **SwapchainKHR** (one-to-many): A presentation request contains an array of target swapchain handles.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:135085` — struct PresentInfoKHR

## PrimRef

Transient builder record for one source primitive and its bounds.

Role: Decouples BVH partitioning from source-geometry storage layouts.

Lifecycle: Generated from geometry during a build, partitioned and consumed by builders, then represented by BVH leaf payloads or discarded with build workspace.

| Field | Type | Meaning |
| --- | --- | --- |
| `bounds` | `BBox3fa` | Axis-aligned bounds used by split heuristics and node construction. |
| `geomID` | `unsigned int accessor` | Source geometry identifier. |
| `primID` | `unsigned int accessor` | Source primitive identifier. |

Relationships:
- PrimRef → **BVHN** (many-to-one): Many primitive references are partitioned into one hierarchy during a BVH build.

Evidence:
- Code: `thirdparty/embree/kernels/builders/primref.h:11` — PrimRef
- Code: `thirdparty/embree/kernels/builders/primrefgen.cpp` — PrimRef construction

## ProjectCatalog

ProjectList-owned catalog, configuration, and selected-project state.

Role: Keeps the project records displayed and acted on by project-manager workflows.

Lifecycle: Created with ProjectList; project scanning populates records, UI actions change selection, and the catalog is discarded with ProjectList.

| Field | Type | Meaning |
| --- | --- | --- |
| `_config_path` | `String` | Path for project-list configuration. |
| `_config` | `ConfigFile` | Configuration backing the project-list state. |
| `_projects` | `Vector<Item>` | Catalogued project records. |
| `_selected_project_paths` | `HashSet<String>` | Keys for the current selected projects. |

Evidence:
- Code: `editor/project_manager/project_list.h:42` — ProjectList

## ProjectSettings

Singleton-backed project configuration store.

Role: Stores and loads project-level properties and related metadata.

Lifecycle: Constructed as the settings singleton, populated by setup or load operations, updated through setting APIs, saved through save operations, and destroyed at shutdown.

| Field | Type | Meaning |
| --- | --- | --- |
| `props` | `RBMap<StringName, VariantContainer>` | Named project properties and their metadata. |
| `resource_path` | `String` | Configured project resource root. |
| `feature_overrides` | `HashMap<StringName, LocalVector<Pair<StringName, StringName>>>` | Feature-specific setting override mappings. |
| `autoloads` | `HashMap<StringName, AutoloadInfo>` | Configured autoload definitions. |
| `_version` | `uint32_t` | Change-tracking version used by cached setting access. |

Evidence:
- Code: `core/config/project_settings.h:40` — ProjectSettings

## PropertyInfo

Reflection description of a named value, including runtime type and editor or usage metadata.

Role: property metadata

Lifecycle: Created while binding or listing properties, consumed by class, object, and script APIs, then released with its containing metadata.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `Variant::Type` | Declared Variant-facing type. |
| `name` | `StringName` | Property name. |
| `hint` | `PropertyHint` | Editor or interpretation hint. |
| `usage` | `uint32_t` | Usage flags. |

Relationships:
- PropertyInfo → **Variant** (one-to-many): PropertyInfo describes Variant-facing values and can be paired with Variant defaults.

Evidence:
- Code: `core/object/property_info.h:124` — PropertyInfo
- Code: `core/object/property_info.cpp` — PropertyInfo::operator==
- Code: `core/object/class_db.cpp:1618` — ClassDB::get_property

## RDAccelerationStructureGeometry

Triangle geometry descriptor for RenderingDevice bottom-level acceleration structures.

Role: Ray-tracing geometry input.

Lifecycle: Created, populated with vertex and optional index-buffer RIDs, passed to blas_create, then released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `vertex_buffer` | `RID` | Buffer containing vertex data. |
| `vertex_count` | `int` | Number of vertices. |
| `index_buffer` | `RID` | Optional buffer containing vertex indices. |
| `index_count` | `int` | Number of indices. |

Relationships:
- RDAccelerationStructureGeometry → **RID** (one-to-many): The geometry descriptor references vertex and optional index buffer IDs.

Evidence:
- Code: `doc/classes/RDAccelerationStructureGeometry.xml:2` — RDAccelerationStructureGeometry

## RDPipelineSpecializationConstant

Selects a value for a shader specialization constant at pipeline creation.

Role: Pipeline variant descriptor.

Lifecycle: Created with a constant identifier and value, attached to an RDPipelineShader, consumed for pipeline creation, and released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `constant_id` | `int` | Per-shader specialization constant identifier. |
| `value` | `Variant` | Value applied to the specialization constant. |

Evidence:
- Code: `doc/classes/RDPipelineSpecializationConstant.xml:2` — RDPipelineSpecializationConstant
- Code: `doc/classes/RDPipelineShader.xml` — RDPipelineShader.specialization_constants

## RDTextureFormat

Describes a RenderingDevice texture's data format, dimensions, usage, and sampling configuration.

Role: GPU texture creation descriptor.

Lifecycle: Constructed and populated before texture creation, consumed by RenderingDevice, then released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `format` | `RenderingDevice.DataFormat` | Pixel data format. |
| `width` | `int` | Texture width. |
| `height` | `int` | Texture height. |
| `usage_bits` | `int` | Allowed texture operations. |

Evidence:
- Code: `doc/classes/RDTextureFormat.xml:2` — RDTextureFormat
- Code: `doc/classes/RDTextureView.xml` — RDTextureView.format_override

## RDUniform

Binds one or more RID-addressed resources to a shader uniform binding.

Role: GPU descriptor binding.

Lifecycle: Created and configured with a binding, type, and IDs, supplied while creating a uniform set, and released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `binding` | `int` | Shader binding index. |
| `uniform_type` | `RenderingDevice.UniformType` | Kind of resource binding. |
| `ids` | `RID[]` | Resources bound to the uniform. |

Relationships:
- RDUniform → **RID** (one-to-many): Each uniform can bind multiple low-level resource IDs.

Evidence:
- Code: `doc/classes/RDUniform.xml:2` — RDUniform

## RID

Opaque, session-only handle used to address a low-level server resource.

Role: Low-level resource identity.

Lifecycle: Obtained from a low-level server or resource, used during the current session, then becomes invalid when its server resource is freed.

Relationships:
- RID → **RDUniform** (one-to-many): An RDUniform can bind an ordered collection of RID values.

Evidence:
- Code: `doc/classes/RID.xml:2` — RID
- Code: `doc/classes/RDUniform.xml` — RDUniform.add_id

## RTCBuffer

Opaque public handle for allocated or shared host-device buffer storage.

Role: Supplies byte-addressed data storage for geometry buffer slots.

Lifecycle: Created with a buffer API call, bound or shared with geometry, and retired through the buffer release API.

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:36` — RTCBuffer
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:39` — rtcNewBuffer
- Code: `thirdparty/embree/include/embree4/rtcore_buffer.h:45` — rtcNewSharedBuffer

## RTCDevice

Opaque public handle for an Embree runtime device.

Role: Owns runtime configuration and is the resource origin for public scene and buffer creation.

Lifecycle: Created with rtcNewDevice, used to create and operate API resources, then retired through the device release API.

Relationships:
- RTCDevice → **RTCScene** (one-to-many): One device can create multiple scene handles.
- RTCDevice → **RTCBuffer** (one-to-many): One device can create multiple buffer handles.

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:11` — RTCDevice
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:15` — rtcNewDevice

## RTCGeometry

Opaque public handle for one typed geometry resource.

Role: Carries geometry configuration, buffer bindings, callbacks, and attachment state.

Lifecycle: Created from a device, configured with buffers and callbacks, attached to a scene and committed, then retired through the geometry release API.

Relationships:
- RTCGeometry → **RTCBuffer** (one-to-many): One geometry can bind multiple buffers in distinct buffer-type slots.

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_geometry.h:15` — RTCGeometry
- Code: `thirdparty/embree/include/embree4/rtcore_geometry.h:169` — rtcSetGeometryBuffer

## RTCHit

Public closest-hit result state.

Role: Reports geometric normal, barycentric coordinates, primitive identity, geometry identity, and instance identity.

Lifecycle: Initialized by the caller as part of a ray-hit query and overwritten when an accepted closer intersection is found.

| Field | Type | Meaning |
| --- | --- | --- |
| `Ng_x` | `float` | X component of the geometric normal. |
| `Ng_y` | `float` | Y component of the geometric normal. |
| `Ng_z` | `float` | Z component of the geometric normal. |
| `u` | `float` | First hit coordinate. |
| `v` | `float` | Second hit coordinate. |
| `primID` | `unsigned int` | Intersected primitive identifier. |
| `geomID` | `unsigned int` | Intersected geometry identifier. |
| `instID` | `unsigned int[RTC_MAX_INSTANCE_LEVEL_COUNT]` | Instance identifier stack. |

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:30` — RTCHit

## RTCPointQuery

Public point-query input structure.

Role: Carries a point, radius, and time to point-query traversal.

Lifecycle: Caller initializes the query, passes it through point-query APIs and callbacks, then reuses or discards caller-owned storage.

| Field | Type | Meaning |
| --- | --- | --- |
| `x` | `float` | X coordinate of the query point. |
| `y` | `float` | Y coordinate of the query point. |
| `z` | `float` | Z coordinate of the query point. |
| `radius` | `float` | Query radius. |
| `time` | `float` | Time used for geometry evaluation. |

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_common.h:364` — RTCPointQuery
- Code: `thirdparty/embree/kernels/common/point_query.h:12` — PointQueryK

## RTCRay

Public scalar ray input state.

Role: Specifies the spatial, interval, time, masking, and identification inputs to a query.

Lifecycle: Initialized by the caller, read and partially updated during traversal, then reused or discarded by the caller.

| Field | Type | Meaning |
| --- | --- | --- |
| `org_x` | `float` | X coordinate of the ray origin. |
| `org_y` | `float` | Y coordinate of the ray origin. |
| `org_z` | `float` | Z coordinate of the ray origin. |
| `tnear` | `float` | Lower parametric query distance. |
| `dir_x` | `float` | X coordinate of the ray direction. |
| `dir_y` | `float` | Y coordinate of the ray direction. |
| `dir_z` | `float` | Z coordinate of the ray direction. |
| `time` | `float` | Time used by motion-blur paths. |
| `tfar` | `float` | Upper parametric distance, tightened by closest-hit traversal. |
| `mask` | `unsigned int` | Ray mask used by query filtering. |
| `id` | `unsigned int` | Caller-provided ray identifier. |
| `flags` | `unsigned int` | Ray-query flags. |

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:11` — RTCRay

## RTCRayHit

Public combined ray and hit structure.

Role: Carries a ray input and mutable closest-hit output through intersection APIs.

Lifecycle: Caller-owned storage is initialized before an intersection call, mutated by accepted intersections, then reused or discarded by the caller.

| Field | Type | Meaning |
| --- | --- | --- |
| `ray` | `RTCRay` | Input ray and current far-distance state. |
| `hit` | `RTCHit` | Closest accepted intersection result. |

Relationships:
- RTCRayHit → **RTCRay** (one-to-one): Each ray-hit record embeds one ray.
- RTCRayHit → **RTCHit** (one-to-one): Each ray-hit record embeds one hit result.

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_ray.h:48` — RTCRayHit

## RTCScene

Opaque public handle for a set of attached geometry and its queryable acceleration structures.

Role: Groups geometry for commit, intersection, occlusion, collision, and point-query operations.

Lifecycle: Created from a device, populated with geometry, committed for queries, and retired through the scene release API.

Relationships:
- RTCScene → **RTCGeometry** (one-to-many): One scene can contain many attached geometry handles.

Evidence:
- Code: `thirdparty/embree/include/embree4/rtcore_device.h:12` — RTCScene
- Code: `thirdparty/embree/include/embree4/rtcore_scene.h:118` — rtcCommitScene

## RVOSimulator2D

Public 2D RVO simulation owner for agents, obstacles, neighbor indexing, and step processing.

Role: coordinates 2D avoidance simulation

Lifecycle: Constructed by the caller, accepts agents and obstacles, computes simulation steps, and is destroyed by the caller.

Evidence:
- Code: `thirdparty/rvo2/rvo2_2d/RVOSimulator2d.h:52` — RVO2D::RVOSimulator2D

## RenderDataRD

Frame render-data object supplied to RD scene rendering.

Role: Aggregates pointers to visible instance and RID collections for renderer consumption.

Lifecycle: Prepared for a scene-rendering operation, passed to RD rendering services, and replaced by data for later rendering work.

| Field | Type | Meaning |
| --- | --- | --- |
| `instances` | `const PagedArray<RenderGeometryInstance *> *` | Points to geometry instances for the render operation. |
| `lights` | `const PagedArray<RID> *` | Points to light IDs. |
| `reflection_probes` | `const PagedArray<RID> *` | Points to reflection-probe IDs. |
| `voxel_gi_instances` | `const PagedArray<RID> *` | Points to voxel-GI instance IDs. |
| `fog_volumes` | `const PagedArray<RID> *` | Points to fog-volume IDs. |

Evidence:
- Code: `servers/rendering/renderer_rd/storage_rd/render_data_rd.h:38` — RenderDataRD

## RenderSceneBuffersConfiguration

Viewport-change configuration used to set up a RenderSceneBuffers object.

Role: Per-viewport rendering buffer configuration.

Lifecycle: Created and populated by the render engine on a viewport change, passed to buffer configuration, and superseded when the viewport changes again.

| Field | Type | Meaning |
| --- | --- | --- |
| `internal_size` | `Vector2i` | Size of the 3D render buffer. |
| `anisotropic_filtering_level` | `RenderingServer.ViewportAnisotropicFiltering` | Anisotropic filtering level. |
| `fsr_sharpness` | `float` | FSR sharpness when FSR upscaling is used. |

Evidence:
- Code: `doc/classes/RenderSceneBuffersConfiguration.xml:2` — RenderSceneBuffersConfiguration
- Code: `doc/classes/RenderSceneBuffers.xml` — RenderSceneBuffers.configure

## RenderingDeviceDriverVulkan::BufferInfo

Cross-cutting Vulkan driver record addressed through buffer driver IDs.

Role: Stores backend buffer state for Vulkan rendering operations.

Lifecycle: Associated with a buffer driver ID, recovered from that ID during driver operations, and retired by the driver's resource-management path; the indexed excerpt does not establish exact destruction behavior.

Evidence:
- Code: `drivers/vulkan/rendering_device_driver_vulkan.h:238` — struct BufferInfo
- Code: `drivers/vulkan/rendering_device_driver_vulkan.cpp` — BufferInfo casts from p_buffer.id

## RenderingDeviceDriverVulkan::CommandBufferInfo

Cross-cutting Vulkan driver record addressed through command-buffer driver IDs.

Role: Stores backend state used while recording or submitting Vulkan command buffers.

Lifecycle: Associated with a command-buffer driver ID and recovered from that ID throughout driver operations; the indexed excerpt does not establish exact destruction behavior.

Evidence:
- Code: `drivers/vulkan/rendering_device_driver_vulkan.h:61` — struct CommandBufferInfo
- Code: `drivers/vulkan/rendering_device_driver_vulkan.cpp` — CommandBufferInfo casts from p_cmd_buffer.id

## Resource

Base persisted asset object used by engine APIs such as textures, themes, tile sets, scripts, shaders, and materials.

Role: Asset data with a loadable and savable engine type.

Lifecycle: Created or imported, loaded by ResourceLoader, optionally localized to a scene, saved by ResourceSaver, and released when no longer referenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `resource_name` | `String` | Human-readable resource name. |
| `resource_path` | `String` | Filesystem path associated with the resource. |
| `resource_local_to_scene` | `bool` | Whether the resource is localized to an instantiated scene. |

Relationships:
- Resource → **SceneState** (many-to-many): A scene state catalogs resources associated with its packed scene.

Evidence:
- Code: `doc/classes/Resource.xml:2` — Resource
- Code: `doc/classes/SceneState.xml:2` — SceneState

## SDL_Event

SDL's externally exchanged union for an event payload and its event type.

Role: moves application, device, and system event data through the event API

Lifecycle: Produced by SDL or the application, copied into or read from the event queue, then discarded or handled by the consumer.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `Uint32` | Common event-type discriminator shared by event alternatives. |
| `payload union` | `SDL_Event` | Alternative typed event payloads such as text, input, drop, and user events. |

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_events.h:983` — SDL_Event
- Code: `thirdparty/sdl/events/SDL_events.c:981` — SDL_AddEvent

## SDL_EventEntry

Internal event-queue record that owns one SDL_Event and any transferred temporary memory.

Role: links queued events into active and reusable lists

Lifecycle: Allocated or reused when an event is queued, removed after consumption or shutdown, and then freed or returned to the free list.

| Field | Type | Meaning |
| --- | --- | --- |
| `event` | `SDL_Event` | Queued event value. |
| `memory` | `SDL_TemporaryMemory *` | Temporary payload allocations associated with the queued event. |
| `prev` | `SDL_EventEntry *` | Previous entry in the doubly linked queue. |
| `next` | `SDL_EventEntry *` | Next entry in the active or free list. |

Relationships:
- SDL_EventEntry → **SDL_Event** (one-to-one): Each queue entry embeds exactly one event value.
- SDL_EventEntry → **SDL_EventEntry** (one-to-one): Each entry can link to one previous and one next queue entry.

Evidence:
- Code: `thirdparty/sdl/events/SDL_events.c:127` — SDL_EventEntry

## SDL_IOStream

Opaque SDL stream handle for file, memory, and application-defined byte I/O.

Role: abstracts stream positioning, reading, writing, closing, and stream metadata

Lifecycle: Created or opened by an SDL stream constructor, used for I/O, then closed through the stream API.

Relationships:
- SDL_IOStream → **SDL_PropertiesID** (one-to-one): Each stream can expose one property group through its I/O-properties API.

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_iostream.h:31` — SDL_IOStream
- Code: `thirdparty/sdl/io/SDL_iostream.c:350` — SDL_GetIOProperties

## SDL_PropertiesID

Opaque identifier for an SDL runtime property group.

Role: shares named runtime metadata and configuration between SDL services

Lifecycle: Created by a properties API, populated or queried by services, and destroyed when the property group is no longer needed.

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_properties.h:28` — SDL_PropertiesID
- Code: `thirdparty/sdl/SDL_properties.c:542` — SDL_GetStringProperty

## SDL_hid_device_info

Linked record describing an enumerated HID device.

Role: exchanges discovered HID metadata between enumeration and device-opening code

Lifecycle: Built during HID enumeration, traversed by consumers, and released by the matching HID free operation.

| Field | Type | Meaning |
| --- | --- | --- |
| `next` | `SDL_hid_device_info *` | Next device-information record in the enumeration list. |

Relationships:
- SDL_hid_device_info → **SDL_hid_device_info** (one-to-one): Each record links to the next record in an enumeration result list.

Evidence:
- Code: `thirdparty/sdl/include/SDL3/SDL_hidapi.h:112` — SDL_hid_device_info
- Code: `thirdparty/sdl/hidapi/SDL_hidapi.c:1410` — SDL_hid_enumerate

## SVG_RendererRec

Renderer extension record that keeps SVG rendering hooks and hook-owned state.

Role: SVG renderer state

Lifecycle: The record exists for its renderer-module lifetime; hooks and state remain available while the renderer is loaded.

| Field | Type | Meaning |
| --- | --- | --- |
| `root` | `FT_RendererRec` | Embedded generic renderer record. |
| `loaded` | `FT_Bool` | Whether SVG renderer state is loaded. |
| `hooks_set` | `FT_Bool` | Whether external SVG hooks have been configured. |
| `hooks` | `SVG_RendererHooks` | External SVG rendering hooks. |
| `state` | `FT_Pointer` | Hook-owned renderer state. |

Evidence:
- Code: `thirdparty/freetype/src/svg/svgtypes.h:27` — SVG_RendererRec_

## Scene Fixture

A persisted `.tscn` resource defining nodes, resources, scripts, animations, and unique-name flags used as editor-test context.

Role: context storage

Lifecycle: Loaded when a completion configuration requests it and retained as shared fixture data.

| Field | Type | Meaning |
| --- | --- | --- |
| `nodes` | `scene node declarations` | The hierarchy and node types exposed to scene-aware completion. |
| `resources` | `subresources and external resources` | Animations or scripts attached to scene elements. |
| `unique_name_in_owner` | `boolean property` | Marks scene nodes addressable through unique-name lookup. |

Evidence:
- Code: `modules/gdscript/tests/scripts/completion/get_node/get_node.tscn` — UniqueAnimationPlayer and UniqueA nodes
- Code: `modules/gdscript/tests/scripts/completion/argument_options/argument_options.tscn:6` — AnimationLibrary_gs7mj

## SceneReplicationConfig

Persisted resource configuring which NodePath properties participate in spawn, synchronization, and change watching.

Role: scene replication policy

Lifecycle: Created or edited as a Resource, saved with its replication properties, read by spawners and synchronizers, and reset or released with the resource.

| Field | Type | Meaning |
| --- | --- | --- |
| `properties` | `List<ReplicationProperty>` | Configured replication-property entries. |
| `spawn_props` | `List<NodePath>` | Derived property paths serialized for spawning. |
| `sync_props` | `List<NodePath>` | Derived property paths synchronized to peers. |
| `watch_props` | `List<NodePath>` | Derived property paths watched for changes. |

Evidence:
- Code: `modules/multiplayer/scene_replication_config.h:36` — SceneReplicationConfig
- Code: `modules/multiplayer/scene_replication_config.cpp:259` — SceneReplicationConfig::get_sync_properties

## SceneState

Serialized tables and records used to reconstruct a packed scene.

Role: Holds shared names, Variant values, node paths, node records, and signal connection records.

Lifecycle: Built while packing a scene, retained by PackedScene, read during instantiation, and released with its resource references.

| Field | Type | Meaning |
| --- | --- | --- |
| `names` | `Vector<StringName>` | Shared serialized names. |
| `variants` | `Vector<Variant>` | Shared serialized values. |
| `node_paths` | `Vector<NodePath>` | Serialized node paths. |
| `nodes` | `Vector<NodeData>` | Serialized node records. |
| `connections` | `Vector<ConnectionData>` | Serialized signal connection records. |

Relationships:
- SceneState → **PackedScene** (one-to-one): A PackedScene retains one SceneState reference.

Evidence:
- Code: `scene/resources/packed_scene.h:38` — SceneState
- Code: `scene/resources/packed_scene.h:167` — SceneState::NodeData
- Code: `scene/resources/packed_scene.h` — SceneState::ConnectionData

## ScriptPathAttribute

Assembly metadata that associates a managed script class with a script file path.

Role: script discovery metadata

Lifecycle: Generated during C# compilation, stored as type metadata in the resulting assembly, read by script discovery, and retired when its assembly unloads.

| Field | Type | Meaning |
| --- | --- | --- |
| `Path` | `string` | The script file path carried by the attribute. |

Evidence:
- Code: `modules/mono/glue/GodotSharp/GodotSharp/Core/Attributes/ScriptPathAttribute.cs:9` — ScriptPathAttribute
- Code: `modules/mono/editor/Godot.NET.Sdk/Godot.SourceGenerators/ScriptPathAttributeGenerator.cs:63` — VisitGodotScriptClass

## Segment

Graphite2 shaped text run containing character information and an ordered slot result.

Role: shaping result

Lifecycle: It is created for text using a face, selected SILF, font, and features; its slots and character information are queried; then it is destroyed by the API client.

| Field | Type | Meaning |
| --- | --- | --- |
| `m_face` | `const Face *` | Face used for glyph lookup and shaping data. |
| `m_silf` | `const Silf *` | Selected Graphite behavior used by the segment. |
| `m_charinfo` | `CharInfo *` | Character-information array exposed through charinfo(index). |

Relationships:
- Segment → **Face** (many-to-one): Each segment refers to the face from which its glyph and behavior data is obtained.

Evidence:
- Code: `thirdparty/graphite/src/inc/Segment.h:213` — graphite2::Segment
- Code: `thirdparty/graphite/src/gr_segment.cpp` — gr_seg_cinfo, gr_seg_first_slot, and gr_seg_last_slot

## ShaderMaterial

Material configuration consisting of a Shader plus values for that shader's parameters.

Role: High-level render material.

Lifecycle: Created or loaded as a Resource, assigned a Shader and parameter values, consumed during rendering, and released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `shader` | `Shader` | Custom shader program used by the material. |

Evidence:
- Code: `doc/classes/ShaderMaterial.xml:2` — ShaderMaterial

## Shape

Polymorphic collision geometry used by bodies and collision queries.

Role: Supplies collision operations, subshape identification, materials, triangle data, and geometry-specific behavior.

Lifecycle: Created from ShapeSettings or concrete shape construction, retained through Ref-based ownership, used by bodies and queries, and released when owning references are released.

Relationships:
- Shape → **PhysicsMaterialSimple** (many-to-one): Concrete shapes can resolve a PhysicsMaterial, falling back to the default material where applicable.

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/Shape.h:36` — Shape
- Code: `thirdparty/jolt_physics/Jolt/Physics/Collision/Shape/ConvexShape.h` — ConvexShape::GetMaterial

## Shape2D

Resource base for 2D physics collision geometry.

Role: Supplies a reusable collision shape to collision owners.

Lifecycle: Created or loaded as a resource, assigned to a collision shape owner, used by physics, then released by reference counting.

Relationships:
- Shape2D → **CollisionShape2D** (one-to-many): One Shape2D resource can be assigned to many CollisionShape2D nodes.

Evidence:
- Code: `doc/classes/CircleShape2D.xml:2` — CircleShape2D
- Code: `doc/classes/CollisionObject2D.xml:2` — CollisionObject2D

## Skeleton

Reference-counted ordered collection of named skeletal joints.

Role: Defines parent relationships used by poses, animation, mapping, and ragdoll construction.

Lifecycle: Built or restored with joints, parent indices are calculated, consumed by poses and ragdolls, and released through reference ownership.

| Field | Type | Meaning |
| --- | --- | --- |
| `mJoints` | `JointVector` | Ordered joints from which parent indices are resolved. |

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/Skeleton.h` — Skeleton and Skeleton::Joint
- Code: `thirdparty/jolt_physics/Jolt/Skeleton/Skeleton.cpp:34` — Skeleton::CalculateParentJointIndices

## SkeletonModificationStack2D

Resource that holds ordered SkeletonModification2D operations.

Role: 2D skeletal-processing plan.

Lifecycle: Created and configured with modifications, assigned to a Skeleton2D, executed in order, and released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `modification_count` | `int` | Number of modifications in the stack. |
| `strength` | `float` | Overall strength with which modifications are applied. |

Evidence:
- Code: `doc/classes/SkeletonModificationStack2D.xml:2` — SkeletonModificationStack2D

## SoftBodySharedSettings

Reference-counted topology and constraint data shared by soft-body instances.

Role: Defines deformable vertices, faces, material assignments, and structural constraints used to initialize and update soft bodies.

Lifecycle: Built or optimized as shared settings, referenced by SoftBodyCreationSettings and SoftBodyMotionProperties, and released when the final reference is released.

| Field | Type | Meaning |
| --- | --- | --- |
| `mVertices` | `Array<Vertex>` | Rest-state soft-body vertices. |
| `mFaces` | `Array<Face>` | Triangle faces connecting vertex indices. |
| `mEdgeConstraints` | `Array<Edge>` | Edge constraints used during deformable-body updates. |
| `mMaterials` | `PhysicsMaterialList` | Materials used by the soft-body faces. |

Relationships:
- SoftBodySharedSettings → **PhysicsMaterialSimple** (many-to-many): Soft-body settings retain a list of materials that can be assigned to faces.

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/SoftBody/SoftBodySharedSettings.h:15` — SoftBodySharedSettings
- Code: `thirdparty/jolt_physics/Jolt/Physics/SoftBody/SoftBodyMotionProperties.h` — SoftBodyMotionProperties::GetSettings, GetMaterials, GetFaces

## SourceLine

A project-upgrade source record that keeps text together with comment classification.

Role: Lets migration passes distinguish source content from comments while transforming project files.

Lifecycle: Created when project source is split into lines, passed through conversion and check methods, then serialized or discarded with the migration operation.

| Field | Type | Meaning |
| --- | --- | --- |
| `line` | `String` | Original or transformed source-line text. |
| `is_comment` | `bool` | Whether the line is classified as a comment. |

Evidence:
- Code: `editor/project_upgrade/project_converter_3_to_4.h:39` — SourceLine
- Code: `editor/project_upgrade/project_converter_3_to_4.h` — ProjectConverter3To4::rename_common

## SpoofImpl

ICU internal spoof-checker implementation.

Role: Holds allowed-character and allowed-locale policy used by spoof-check APIs.

Lifecycle: Created behind a checker handle, validated before operations, and released with the checker implementation.

| Field | Type | Meaning |
| --- | --- | --- |
| `fAllowedCharsSet` | `const UnicodeSet *` | References allowed Unicode characters. |
| `fAllowedLocales` | `const char *` | Stores the allowed-locale list. |

Evidence:
- Code: `thirdparty/icu4c/i18n/uspoof_impl.h:55` — SpoofImpl
- Code: `thirdparty/icu4c/i18n/uspoof_impl.cpp:110` — SpoofImpl::validateThis

## SpvReflectBlockVariable

Reflected block or push-constant variable record.

Role: Carries byte offsets, sizes, decorations, numeric/array traits, and nested member metadata.

Lifecycle: Owned by reflection-module data; push-constant enumeration documentation says callers must not free returned block pointers.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `const char*` | Block-variable name. |
| `offset` | `uint32_t` | Byte offset. |
| `absolute_offset` | `uint32_t` | Absolute byte offset. |
| `size` | `uint32_t` | Byte size. |
| `padded_size` | `uint32_t` | Padded byte size. |
| `members` | `SpvReflectBlockVariable*` | Nested block-variable members. |
| `type_description` | `SpvReflectTypeDescription*` | Underlying reflected type description. |

Relationships:
- SpvReflectBlockVariable → **SpvReflectBlockVariable** (zero-to-many): member_count and members describe nested block-variable members.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:457` — SpvReflectBlockVariable
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:982` — spvReflectEnumeratePushConstantBlocks

## SpvReflectDescriptorBinding

Reflected descriptor-binding record.

Role: Describes one shader resource binding and its set, type, array, block, and optional counter binding.

Lifecycle: Owned by reflection-module data; pointer-returning APIs state that callers must not free the returned binding pointer.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `const char*` | Binding name. |
| `binding` | `uint32_t` | Binding number. |
| `set` | `uint32_t` | Descriptor-set number. |
| `descriptor_type` | `SpvReflectDescriptorType` | Reflected descriptor type. |
| `resource_type` | `SpvReflectResourceType` | Reflected resource category. |
| `block` | `SpvReflectBlockVariable` | Associated block metadata. |
| `array` | `SpvReflectBindingArrayTraits` | Binding-array dimensions. |
| `count` | `uint32_t` | Descriptor count. |

Relationships:
- SpvReflectDescriptorBinding → **SpvReflectBlockVariable** (zero-to-one): A descriptor binding embeds one block metadata record for block-backed resources.
- SpvReflectDescriptorBinding → **SpvReflectDescriptorBinding** (zero-to-one): uav_counter_binding can point to one counter binding.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:484` — SpvReflectDescriptorBinding
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:1061` — spvReflectGetDescriptorBinding

## SpvReflectDescriptorSet

Reflected descriptor-set record.

Role: Groups descriptor-binding pointers by set number.

Lifecycle: Owned by its containing reflection module or entry-point reflection data; callers are instructed not to free enumerated set pointers.

| Field | Type | Meaning |
| --- | --- | --- |
| `set` | `uint32_t` | Descriptor-set number. |
| `binding_count` | `uint32_t` | Number of bindings in the set. |
| `bindings` | `SpvReflectDescriptorBinding**` | Pointers to the set's descriptor bindings. |

Relationships:
- SpvReflectDescriptorSet → **SpvReflectDescriptorBinding** (one-to-many): A descriptor set stores binding_count and a bindings pointer array.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:517` — SpvReflectDescriptorSet
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:774` — spvReflectEnumerateDescriptorSets

## SpvReflectEntryPoint

Reflected shader entry-point record.

Role: Groups shader stage, interface variables, descriptor sets, push-constant use, and execution-mode data for one entry point.

Lifecycle: Owned by the containing SpvReflectShaderModule and valid while that module remains valid.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `const char*` | Entry-point name. |
| `id` | `uint32_t` | SPIR-V identifier for the entry point. |
| `shader_stage` | `SpvReflectShaderStageFlagBits` | Reflected shader stage. |
| `input_variable_count` | `uint32_t` | Number of input interface variables. |
| `output_variable_count` | `uint32_t` | Number of output interface variables. |
| `descriptor_set_count` | `uint32_t` | Number of descriptor sets used by the entry point. |
| `local_size` | `struct LocalSize` | Three-dimensional local size record. |

Relationships:
- SpvReflectEntryPoint → **SpvReflectInterfaceVariable** (one-to-many): The entry point stores input_variables, output_variables, and interface_variables collections.
- SpvReflectEntryPoint → **SpvReflectDescriptorSet** (one-to-many): The entry point stores descriptor_set_count and descriptor_sets.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:530` — SpvReflectEntryPoint

## SpvReflectInterfaceVariable

Reflected input or output SPIR-V variable record.

Role: Carries interface location, storage class, semantic, numeric/array traits, format, and nested member metadata.

Lifecycle: Owned by reflection-module data; enumeration documentation says callers must not free returned interface-variable pointers.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `const char*` | Variable name. |
| `location` | `uint32_t` | Interface location. |
| `storage_class` | `SpvStorageClass` | SPIR-V storage class. |
| `semantic` | `const char*` | Optional semantic string. |
| `numeric` | `SpvReflectNumericTraits` | Numeric type traits. |
| `array` | `SpvReflectArrayTraits` | Array traits. |
| `format` | `SpvReflectFormat` | Reflected format. |
| `members` | `SpvReflectInterfaceVariable*` | Nested interface-variable members. |

Relationships:
- SpvReflectInterfaceVariable → **SpvReflectInterfaceVariable** (zero-to-many): member_count and members describe nested interface-variable members.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:424` — SpvReflectInterfaceVariable
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:826` — spvReflectEnumerateInterfaceVariables

## SpvReflectShaderModule

Top-level reflection record for compiled SPIR-V code and its reflected metadata.

Role: Owns the module-level collections returned by the reflection API.

Lifecycle: A caller supplies the record to spvReflectCreateShaderModule or spvReflectCreateShaderModule2 and retires it with spvReflectDestroyShaderModule.

| Field | Type | Meaning |
| --- | --- | --- |
| `entry_point_count` | `uint32_t` | Number of reflected entry points. |
| `entry_points` | `SpvReflectEntryPoint*` | Array of reflected entry-point records. |
| `descriptor_binding_count` | `uint32_t` | Number of descriptor bindings associated with the first entry point. |
| `descriptor_bindings` | `SpvReflectDescriptorBinding*` | Reflected descriptor-binding records. |
| `descriptor_set_count` | `uint32_t` | Number of reflected descriptor sets associated with the first entry point. |
| `interface_variables` | `SpvReflectInterfaceVariable*` | Reflected interface-variable records. |
| `push_constant_blocks` | `SpvReflectBlockVariable*` | Reflected push-constant block records. |
| `spec_constants` | `SpvReflectSpecializationConstant*` | Reflected specialization-constant records. |

Relationships:
- SpvReflectShaderModule → **SpvReflectEntryPoint** (one-to-many): The module stores entry_point_count and an entry_points array.
- SpvReflectShaderModule → **SpvReflectDescriptorBinding** (one-to-many): The module stores descriptor_binding_count and descriptor_bindings.
- SpvReflectShaderModule → **SpvReflectDescriptorSet** (one-to-many): The module stores descriptor_set_count and descriptor_sets.
- SpvReflectShaderModule → **SpvReflectInterfaceVariable** (one-to-many): The module stores interface-variable counts and interface_variables.
- SpvReflectShaderModule → **SpvReflectBlockVariable** (one-to-many): The module stores push_constant_block_count and push_constant_blocks.
- SpvReflectShaderModule → **SpvReflectSpecializationConstant** (one-to-many): The module stores spec_constant_count and spec_constants.

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:600` — SpvReflectShaderModule
- Code: `thirdparty/spirv-reflect/spirv_reflect.h` — spvReflectCreateShaderModule, spvReflectDestroyShaderModule

## SpvReflectSpecializationConstant

Reflected shader specialization-constant record.

Role: Carries the SPIR-V id, constant id, name, type description, and raw default-value data.

Lifecycle: Owned by reflection-module data; specialization-constant enumeration documentation says callers must not free returned pointers.

| Field | Type | Meaning |
| --- | --- | --- |
| `spirv_id` | `uint32_t` | SPIR-V identifier. |
| `constant_id` | `uint32_t` | Specialization constant identifier. |
| `name` | `const char*` | Constant name. |
| `type_description` | `SpvReflectTypeDescription*` | Reflected type description. |
| `default_value_size` | `uint32_t` | Size in bytes of the raw default value. |
| `default_value` | `void*` | Raw default-value data. |

Evidence:
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:579` — SpvReflectSpecializationConstant
- Code: `thirdparty/spirv-reflect/spirv_reflect.h:1042` — spvReflectEnumerateSpecializationConstants

## StreamPeerBuffer

In-memory binary stream backed by a byte array and cursor.

Role: Binary data exchange buffer.

Lifecycle: Created, written or initialized with data, read through its moving cursor, cleared or duplicated, then released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `data_array` | `PackedByteArray` | Underlying byte-array stream data. |

Evidence:
- Code: `doc/classes/StreamPeerBuffer.xml:2` — StreamPeerBuffer

## StringName

Interned string-name value used for class, property, method, signal, path, and translation identifiers.

Role: canonical identifier

Lifecycle: Interned into StringName table data when constructed, shared by value copies, and released as references leave scope.

| Field | Type | Meaning |
| --- | --- | --- |
| `_data` | `_Data *` | Pointer to interned name data. |

Relationships:
- StringName → **NodePath** (one-to-many): NodePath stores path and subpath segments as StringName values.
- StringName → **Translation** (one-to-many): Translation uses StringName values for message keys and translated strings.

Evidence:
- Code: `core/string/string_name.h:40` — StringName
- Code: `core/string/string_name.h` — StringName::_Data
- Code: `core/string/string_name.cpp:38` — StringName::Table

## SubGrid

Embree primitive identifying a local grid region and its source geometry.

Role: Connects subgrid intersection code to a GridMesh grid and its vertex data.

Lifecycle: Created with x, y, geometry, and primitive identifiers; queried by subgrid intersection kernels; retired with the owning primitive collection.

| Field | Type | Meaning |
| --- | --- | --- |
| `x` | `unsigned int` | Horizontal subgrid coordinate. |
| `y` | `unsigned int` | Vertical subgrid coordinate. |
| `geomID` | `unsigned int` | Identifier used to retrieve the source GridMesh from the scene. |
| `primID` | `unsigned int` | Identifier used to retrieve the GridMesh::Grid record. |

Evidence:
- Code: `thirdparty/embree/kernels/geometry/subgrid.h:13` — SubGrid
- Code: `thirdparty/embree/kernels/geometry/subgrid_intersector.h:21` — SubGridIntersector1Moeller

## SwapchainKHR

Opaque Vulkan swapchain handle referenced by presentation requests.

Role: identifies a presentation target

Lifecycle: A `SwapchainCreateInfoKHR` configures creation and `PresentInfoKHR` references the handle; handle retirement is outside the supplied declarations.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:154367` — struct SwapchainCreateInfoKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:135085` — struct PresentInfoKHR

## Test Script Fixture

A persisted GDScript source case that supplies code for one parser, analyzer, runtime, completion, or LSP scenario.

Role: test input

Lifecycle: Created as a categorized `.gd` fixture, consumed by its relevant test path, and retained with the suite as the regression case.

| Field | Type | Meaning |
| --- | --- | --- |
| `source` | `GDScript text` | The program or cursor-bearing editor input under test. |
| `category` | `path segment` | The validation stage represented by its directory, such as `parser`, `analyzer`, or `runtime`. |
| `test entry point` | `function location` | Usually the recognizable `test()` or editor callback location that exercises the case. |

Relationships:
- Test Script Fixture → **Expected Result Fixture** (one-to-one): One executable fixture conventionally has one same-basename `.out` expectation.

Evidence:
- Code: `modules/gdscript/tests/scripts/analyzer/features/lambda_typed.gd:1` — test()
- Code: `modules/gdscript/tests/scripts/completion/common/identifiers_in_call.gd` — _init()

## Theme

Resource containing reusable visual settings for Control and Window nodes.

Role: UI styling data.

Lifecycle: Created or loaded, assigned at project or Control scope, inherited by eligible child controls, and released when unreferenced.

Evidence:
- Code: `doc/classes/Theme.xml:2` — Theme
- Code: `doc/classes/StyleBox.xml:2` — StyleBox

## Theora Stream Information

Persisted stream-header metadata for a Theora video stream.

Role: Defines coded-frame geometry, visible picture geometry, timing, and pixel format for encoder and decoder setup.

Lifecycle: Initialized before header parsing or encoding setup, populated from Theora headers or application configuration, then cleared after codec-context setup is complete.

| Field | Type | Meaning |
| --- | --- | --- |
| `frame_width` | `ogg_uint32_t` | Encoded frame width. |
| `frame_height` | `ogg_uint32_t` | Encoded frame height. |
| `pic_width` | `ogg_uint32_t` | Visible picture width. |
| `pic_height` | `ogg_uint32_t` | Visible picture height. |
| `fps_numerator` | `ogg_uint32_t` | Frame-rate numerator. |
| `fps_denominator` | `ogg_uint32_t` | Frame-rate denominator. |
| `pixel_fmt` | `th_pixel_fmt` | Chroma pixel format. |

Evidence:
- Code: `thirdparty/libtheora/theora/codec.h:157` — th_info

## TileData

Per-tile settings exposed from a TileSet source.

Role: Tile behavior and rendering data.

Lifecycle: Created and edited through a TileSet source, queried by tile-map code, optionally updated at runtime, and retained with its source.

Relationships:
- TileData → **TileSet** (many-to-one): Each TileData object represents a tile in one TileSet source.

Evidence:
- Code: `doc/classes/TileData.xml:2` — TileData

## TileMapLayerCellData

Per-coordinate cell state retained by TileMapLayer.

Role: Connects a tile-map coordinate with its TileMapCell and supports derived quadrants.

Lifecycle: Created or updated when a layer cell changes and removed when the cell is erased or the layer ends.

| Field | Type | Meaning |
| --- | --- | --- |
| `coords` | `Vector2i` | Grid coordinate of the cell. |
| `cell` | `TileMapCell` | Tile selection and alternative data for the coordinate. |

Evidence:
- Code: `scene/2d/tile_map_layer.h:105` — CellData
- Code: `scene/2d/tile_map_layer.h` — TileMapLayer::get_tile_map_layer_data
- Code: `scene/2d/tile_map_layer.cpp` — CellData::coords and CellData::cell usage

## TileSet

Resource library of tile sources and tile property layers for TileMapLayer.

Role: Reusable tile catalog.

Lifecycle: Created or imported, populated with sources and layers, assigned to tile maps, serialized as a Resource, and released when unreferenced.

| Field | Type | Meaning |
| --- | --- | --- |
| `tile_size` | `Vector2i` | Base tile dimensions. |

Relationships:
- TileSet → **TileData** (one-to-many): Tile sources exposed through a TileSet provide data for many individual tiles.

Evidence:
- Code: `doc/classes/TileSet.xml:2` — TileSet
- Code: `doc/classes/TileData.xml:2` — TileData

## Transform3D

Three-dimensional transform value composed of a Basis and an origin vector.

Role: spatial transform value

Lifecycle: Constructed as a value, composed or interpolated in math code, transported through dynamic conversion where needed, and discarded as a value object.

| Field | Type | Meaning |
| --- | --- | --- |
| `basis` | `Basis` | Linear transform basis. |
| `origin` | `Vector3` | Translation component. |

Evidence:
- Code: `core/math/transform_3d.h:38` — Transform3D
- Code: `core/math/transform_3d.cpp:96` — Transform3D::interpolate_with
- Code: `core/variant/method_ptrcall.h:203` — PtrToArg<Transform3D>

## Translation

Resource-like translation catalog mapping contextual message keys to translated StringName vectors.

Role: localized message catalog

Lifecycle: Configured with locale and messages, attached to a translation domain, queried for singular or plural messages, and released through reference management.

| Field | Type | Meaning |
| --- | --- | --- |
| `locale` | `String` | Locale associated with the catalog. |
| `translation_map` | `HashMap<MessageKey, Vector<StringName>>` | Contextual source-message to translated-message mapping. |

Relationships:
- Translation → **TranslationDomain** (one-to-many): A TranslationDomain can select from multiple Translation catalogs.
- Translation → **StringName** (one-to-many): Message keys and translated strings use interned names.

Evidence:
- Code: `core/string/translation.h:38` — Translation
- Code: `core/string/translation.h` — Translation::MessageKey
- Code: `core/string/translation.cpp:164` — Translation::get_message

## TranslationDomain

Reference-counted selector for translation catalogs, locale override, fallback behavior, and pseudolocalization configuration.

Role: translation resolution scope

Lifecycle: Created or retrieved by TranslationServer, populated with Translation references, queried for messages, and released when references end.

| Field | Type | Meaning |
| --- | --- | --- |
| `locale_override` | `String` | Optional locale used instead of the server locale. |
| `translation_set` | `HashSet<Ref<Translation>>` | Catalogs available to the domain. |

Relationships:
- TranslationDomain → **Translation** (one-to-many): A TranslationDomain holds a set of Translation references.

Evidence:
- Code: `core/string/translation_domain.h:37` — TranslationDomain
- Code: `core/string/translation_domain.cpp:289` — TranslationDomain::get_translations
- Code: `core/string/translation_server.cpp:581` — TranslationServer::get_or_add_domain

## TriRef

Per-triangle provenance reference retained by Manifold mesh relations.

Role: source-triangle provenance

Lifecycle: Created or remapped during boolean and mesh operations, carried in MeshRelationD, and used while result properties are transferred.

| Field | Type | Meaning |
| --- | --- | --- |
| `meshID` | `int` | Identifier of the referenced source mesh. |
| `faceID` | `int` | Referenced source triangle or face identifier. |

Evidence:
- Code: `thirdparty/manifold/src/shared.h:135` — struct TriRef
- Code: `thirdparty/manifold/src/impl.h` — struct Manifold::Impl::MeshRelationD
- Code: `thirdparty/manifold/src/boolean_result.cpp` — struct MapTriRef; outR.meshRelation_.triRef

## TriangleMesh

Reference-counted triangle mesh storing vertices, triangles, and an internal BVH.

Role: accelerated mesh query data

Lifecycle: Built from mesh-face data, queried through its triangle and BVH data, and retired when its reference count reaches zero.

| Field | Type | Meaning |
| --- | --- | --- |
| `triangles` | `Vector<Triangle>` | Triangle records. |
| `vertices` | `Vector<Vector3>` | Mesh vertex positions. |
| `bvh` | `Vector<BVH>` | Triangle acceleration hierarchy. |

Relationships:
- TriangleMesh → **AABB** (one-to-many): TriangleMesh BVH data partitions triangle bounds for spatial queries.

Evidence:
- Code: `core/math/triangle_mesh.h:36` — TriangleMesh
- Code: `core/math/triangle_mesh.h` — TriangleMesh::Triangle
- Code: `core/math/triangle_mesh.h` — TriangleMesh::BVH
- Code: `core/math/triangle_mesh.cpp:108` — TriangleMesh::create

## TwoBodyConstraint

Runtime constraint connecting two bodies.

Role: Participates in island building, body activation, and solver processing for a pair of bodies.

Lifecycle: Created by a concrete TwoBodyConstraintSettings subtype, attached to two bodies, included in simulation islands while active, and released by constraint management.

| Field | Type | Meaning |
| --- | --- | --- |
| `mBody1` | `Body *` | First constrained body used during island construction. |
| `mBody2` | `Body *` | Second constrained body used during island construction. |

Evidence:
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/TwoBodyConstraint.cpp:23` — TwoBodyConstraint::BuildIslands
- Code: `thirdparty/jolt_physics/Jolt/Physics/Constraints/TwoBodyConstraint.h:12` — TwoBodyConstraint

## UCPTrie

ICU immutable Unicode code-point trie.

Role: Maps Unicode code points to compact integer-backed property data.

Lifecycle: Opened from or built into binary trie data, queried by Unicode services, and closed or retained with its owning data set.

| Field | Type | Meaning |
| --- | --- | --- |
| `index` | `const uint16_t *` | References compact trie index units. |
| `name` | `const char *` | Stores an identifying name. |

Evidence:
- Code: `thirdparty/icu4c/common/unicode/ucptrie.h:24` — UCPTrie
- Code: `thirdparty/icu4c/common/normalizer2impl.h` — Normalizer2Impl::normTrie

## UConverter

ICU converter instance for character-encoding conversion.

Role: Maintains conversion state and delegates conversion behavior to shared implementation data.

Lifecycle: Opened or created for conversion, used on source and target buffers, then closed by the converter API owner.

| Field | Type | Meaning |
| --- | --- | --- |
| `sharedData` | `UConverterSharedData *` | References converter metadata and implementation functions. |
| `fromUContext` | `const void *` | Stores context for conversion from Unicode. |
| `toUContext` | `const void *` | Stores context for conversion to Unicode. |

Relationships:
- UConverter → **UConverterSharedData** (many-to-one): Many converter instances can share one converter-data record.

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h:34` — UConverter
- Code: `thirdparty/icu4c/common/ucnv.cpp:384` — converter->sharedData

## UConverterSharedData

Shared immutable or long-lived data for an ICU converter implementation.

Role: Connects converter state to static conversion data and an implementation dispatch record.

Lifecycle: Loaded or constructed as shared converter data, referenced by converters, and cleaned up with converter-data ownership.

| Field | Type | Meaning |
| --- | --- | --- |
| `dataMemory` | `const void *` | References loaded data memory. |
| `staticData` | `const UConverterStaticData *` | References non-changing converter metadata. |
| `impl` | `const UConverterImpl *` | References the vtable-style implementation record. |

Relationships:
- UConverterSharedData → **UConverter** (one-to-many): Shared converter data can be referenced by converter instances.

Evidence:
- Code: `thirdparty/icu4c/common/ucnv_bld.h:93` — UConverterSharedData

## UPNPDev

A discovered UPnP device node.

Role: Discovery result transported from SSDP discovery to IGD selection and description parsing.

Lifecycle: Discovery creates a linked device list; callers traverse it for IGD selection and release the list through the MiniUPnPc cleanup path.

| Field | Type | Meaning |
| --- | --- | --- |
| `pNext` | `struct UPNPDev *` | Links this device node to the next discovered device. |

Relationships:
- UPNPDev → **UPNPDev** (one-to-many): One list head can reach many device nodes through successive `pNext` links.

Evidence:
- Code: `thirdparty/miniupnpc/include/miniupnpc/upnpdev.h:24` — struct UPNPDev
- Code: `thirdparty/miniupnpc/src/miniupnpc.c:505` — UPNP_GetValidIGD

## UPNPDevice

RefCounted UPnP device object.

Role: Represents a discovered device for device-level UPnP operations.

Lifecycle: A UPNP implementation discovers or constructs device objects, callers perform device operations, and normal reference counting retires them.

Evidence:
- Code: `modules/upnp/upnp_device.h:36` — class UPNPDevice : public RefCounted
- Code: `modules/upnp/doc_classes/UPNPDevice.xml` — class UPNPDevice

## UResourceBundle

C-facing ICU resource-bundle state.

Role: Carries the key and resource data needed to retrieve locale-scoped resources.

Lifecycle: Opened against resource data, queried for values, and closed by resource-bundle callers.

| Field | Type | Meaning |
| --- | --- | --- |
| `fKey` | `const char *` | Stores the current resource key. |

Evidence:
- Code: `thirdparty/icu4c/common/uresimp.h:63` — UResourceBundle
- Code: `thirdparty/icu4c/common/uresdata.h:513` — ResourceDataValue

## VP8Encoder

Internal lossy-encoder session state.

Role: Coordinates configuration, picture input, macroblock state, probabilities, partitions, quantization, and analysis state.

Lifecycle: Allocate and initialize for one lossy encode, process macroblocks and partitions, emit the bitstream, release encoder-owned working memory.

| Field | Type | Meaning |
| --- | --- | --- |
| `config` | `const WebPConfig*` | User configuration for the encode session. |
| `pic` | `WebPPicture*` | Input and reporting picture. |
| `mb_w/mb_h` | `int` | Macroblock grid dimensions. |
| `mb_info` | `VP8MBInfo*` | Per-macroblock analysis and segmentation information. |
| `dqm` | `VP8SegmentInfo[]` | Segment quantization and filtering parameters. |
| `proba` | `VP8EncProba` | Coding probabilities and statistics. |
| `parts` | `VP8BitWriter*` | Token-partition writers. |
| `preds` | `uint8_t*` | Prediction-mode map. |

Relationships:
- VP8Encoder → **WebPConfig** (many-to-one): An internal encoder session reads one caller configuration.
- VP8Encoder → **WebPPicture** (many-to-one): An internal encoder session processes one caller picture.

Evidence:
- Code: `thirdparty/libwebp/src/enc/vp8i_enc.h:139` — struct VP8Encoder
- Code: `thirdparty/libwebp/src/enc/webp_enc.c` — InitVP8Encoder, WebPEncode
- Code: `thirdparty/libwebp/src/enc/frame_enc.c:788` — VP8EncTokenLoop

## VP8LHashChain

Lossless match-finding table over image positions.

Role: Stores candidate copy length and offset information used to build backward references.

Lifecycle: Allocate for image dimensions and chosen window settings, populate while scanning pixels, use during reference construction, then delete with lossless encoder scratch state.

| Field | Type | Meaning |
| --- | --- | --- |
| `offset_length_` | `uint32_t*` | Packed copy-offset and match-length entries. |
| `size_` | `int` | Number of image positions represented. |

Evidence:
- Code: `thirdparty/libwebp/src/enc/backward_references_enc.h:123` — struct VP8LHashChain
- Code: `thirdparty/libwebp/src/enc/backward_references_enc.c` — VP8LHashChainFindCopy, VP8LHashChainFindOffset

## Variant

Tagged dynamic runtime value used at object, call, container, parser, and binding boundaries.

Role: runtime value interchange

Lifecycle: Constructed with an active type, copied or converted through Variant operations, and destroyed through type-specific destruction support.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `Variant::Type` | Active runtime type tag. |
| `_data` | `internal payload storage` | Storage accessed by type-specific Variant implementation paths. |

Relationships:
- Variant → **Array** (zero-to-one): A Variant may hold one Array value.
- Variant → **Dictionary** (zero-to-one): A Variant may hold one Dictionary value.
- Variant → **Object** (zero-to-one): A Variant can represent an engine object reference.

Evidence:
- Code: `core/variant/variant.h:94` — Variant
- Code: `core/variant/variant_internal.h:577` — VariantInternalAccessor
- Code: `core/variant/variant_destruct.h:37` — VariantDestruct

## VideoSessionCreateInfoKHR

Externally exchanged configuration record for a Vulkan video session.

Role: selects the video profile and standard-header version for a session

Lifecycle: A caller populates the configuration before a video-session creation operation; the supplied partition declares no session-retirement operation.

| Field | Type | Meaning |
| --- | --- | --- |
| `pVideoProfile` | `const VideoProfileInfoKHR *` | Pointer to the selected video profile. |
| `pStdHeaderVersion` | `const ExtensionProperties *` | Pointer to standard-header version information. |
| `pNext` | `const void *` | Optional extension-chain pointer. |

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_structs.hpp:172949` — struct VideoSessionCreateInfoKHR

## VideoStreamTheora

VideoStream resource for Theora playback.

Role: Stores a Theora video stream consumed by a matching playback implementation.

Lifecycle: A resource loader constructs the stream resource; playback consumes it through VideoStreamPlaybackTheora; editor writing is handled by MovieWriterOGV.

Evidence:
- Code: `modules/theora/video_stream_theora.h` — class VideoStreamTheora : public VideoStream; class VideoStreamPlaybackTheora
- Code: `modules/theora/editor/movie_writer_ogv.h:40` — class MovieWriterOGV : public MovieWriter

## VisualShaderConnection

A graph edge that identifies source and destination node ports.

Role: Connects shader-node outputs to shader-node inputs.

Lifecycle: Created or removed with a VisualShader graph edit and retained while present in that graph.

| Field | Type | Meaning |
| --- | --- | --- |
| `from_node` | `int` | Source node identifier. |
| `from_port` | `int` | Source output-port index. |
| `to_node` | `int` | Destination node identifier. |
| `to_port` | `int` | Destination input-port index. |

Evidence:
- Code: `modules/visual_shader/visual_shader.h` — VisualShader::Connection
- Code: `modules/visual_shader/visual_shader.cpp:1063` — VisualShader::Connection

## VkAndroidHardwareBufferFormatPropertiesANDROID

Optional extension result that carries Vulkan format and suggested YCbCr conversion properties for an Android hardware buffer.

Role: external-image format properties

Lifecycle: The caller chains it from VkAndroidHardwareBufferPropertiesANDROID before the hardware-buffer properties query.

| Field | Type | Meaning |
| --- | --- | --- |
| `format` | `VkFormat` | Reports the Vulkan format. |
| `externalFormat` | `uint64_t` | Reports the external format identifier. |
| `formatFeatures` | `VkFormatFeatureFlags` | Reports supported format features. |
| `samplerYcbcrConversionComponents` | `VkComponentMapping` | Reports suggested component mapping for YCbCr conversion. |

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_android.h:66` — VkAndroidHardwareBufferFormatPropertiesANDROID
- External (official, verified): [Vulkan Specification: Memory Allocation](https://docs.vulkan.org/spec/latest/chapters/memory.html), accessed 2026-07-16

## VkSurfaceKHR

Opaque Vulkan presentation-surface handle returned through the Android surface creation command.

Role: presentation handle

Lifecycle: Written to the command's output pointer on creation; retirement behavior is outside the inspected Android platform header.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_android.h:39` — vkCreateAndroidSurfaceKHR
- Code: `thirdparty/vulkan/include/vulkan/vulkan_handles.hpp:2330` — SurfaceKHR

## VkViSurfaceCreateInfoNN

Externally exchanged input record for creating a Nintendo VI Vulkan surface.

Role: supplies platform window and extension-chain information

Lifecycle: A caller initializes the record and passes it by pointer to `vkCreateViSurfaceNN`; the call produces a `VkSurfaceKHR` through its output parameter.

| Field | Type | Meaning |
| --- | --- | --- |
| `sType` | `VkStructureType` | Record type discriminator. |
| `pNext` | `const void *` | Optional extension-chain pointer. |
| `flags` | `VkViSurfaceCreateFlagsNN` | VI surface creation flags. |
| `window` | `void *` | Native Nintendo VI window handle. |

Relationships:
- VkViSurfaceCreateInfoNN → **VkSurfaceKHR** (zero-to-one): One successful surface-creation call writes one surface handle for the supplied creation record.

Evidence:
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:27` — VkViSurfaceCreateInfoNN
- Code: `thirdparty/vulkan/include/vulkan/vulkan_vi.h:38` — vkCreateViSurfaceNN

## VmaAllocator

Opaque VMA handle that owns Vulkan memory-allocation management state.

Role: allocation-management context

Lifecycle: It is returned by allocator creation and is then passed to VMA allocation, pool, statistics, and memory-management operations.

Evidence:
- Code: `thirdparty/vulkan/vk_mem_alloc.h` — VmaAllocator, VmaAllocator_T, vmaCreateAllocator
- Code: `thirdparty/vulkan/patches/0003-VMA-add-vmaCalculateLazilyAllocatedBytes.patch:19` — vmaCalculateLazilyAllocatedBytes

## VmaAllocatorCreateInfo

Configuration structure supplied when creating a VMA allocator.

Role: Vulkan memory allocator creation input

Lifecycle: The caller prepares this input for vmaCreateAllocator; the created VmaAllocator represents the subsequent allocation-management lifetime.

| Field | Type | Meaning |
| --- | --- | --- |
| `pAllocationCallbacks` | `const VkAllocationCallbacks *` | Optional Vulkan allocation-callback configuration. |
| `pDeviceMemoryCallbacks` | `const VmaDeviceMemoryCallbacks *` | Optional device-memory callback configuration. |
| `pHeapSizeLimit` | `const VkDeviceSize *` | Optional per-heap size limits. |
| `pVulkanFunctions` | `const VmaVulkanFunctions *` | Optional Vulkan function table. |
| `pTypeExternalMemoryHandleTypes` | `const VkExternalMemoryHandleTypeFlagsKHR *` | Optional external-memory handle-type information by memory type. |

Relationships:
- VmaAllocatorCreateInfo → **VmaAllocator** (zero-or-one): One creation configuration can produce one allocator handle when vmaCreateAllocator succeeds.

Evidence:
- Code: `thirdparty/vulkan/vk_mem_alloc.h:355` — VmaAllocatorCreateInfo
- Code: `thirdparty/vulkan/vk_mem_alloc.h:878` — vmaCreateAllocator

## Vorbis Block

A codec work block holding PCM, bit-packing state, window selection, and allocation-chain state.

Role: Carries one unit of Vorbis analysis or synthesis work.

Lifecycle: Initialized against a DSP state, filled during analysis or synthesis, converted to or from an Ogg packet, and cleared or reused by the codec.

| Field | Type | Meaning |
| --- | --- | --- |
| `pcm` | `float **` | Per-channel PCM samples. |
| `opb` | `oggpack_buffer` | Ogg bit-packing buffer for the block. |
| `lW` | `long` | Previous window flag. |
| `W` | `long` | Current window flag. |
| `nW` | `long` | Next window flag. |
| `vd` | `vorbis_dsp_state *` | Owning DSP state. |

Relationships:
- Vorbis Block → **Ogg Packet** (one-to-one): A synthesis call consumes one Ogg packet to populate one Vorbis block.

Evidence:
- Code: `thirdparty/libvorbis/vorbis/codec.h:87` — vorbis_block
- Code: `thirdparty/libvorbis/synthesis.c:25` — vorbis_synthesis

## WaylandThread.WindowState

Cross-cutting state structure for a Wayland window.

Role: retains Wayland protocol objects and frame-callback state for a tracked window

Lifecycle: Held by the Wayland thread while a window is tracked; teardown details are not exposed by the supplied index.

| Field | Type | Meaning |
| --- | --- | --- |
| `frame_callback` | `wl_callback *` | Wayland frame-callback pointer. |
| `wl_surface` | `wl_surface *` | Wayland surface pointer. |
| `xdg_surface` | `xdg_surface *` | XDG surface pointer. |
| `xdg_toplevel` | `xdg_toplevel *` | XDG toplevel pointer. |

Evidence:
- Code: `platform/linuxbsd/wayland/wayland_thread.h` — WaylandThread::WindowState
- Code: `platform/linuxbsd/wayland/wayland_thread.cpp:4612` — WaylandThread::window_get_state

## WebP Decoder State

Internal VP8 WebP decoder state, including errors and optional alpha-decoding data.

Role: Coordinates bitstream decode, output, and alpha-plane handling.

Lifecycle: Allocated for a WebP decode, filled while headers and partitions are parsed, emits rows through I/O callbacks, and is released after decoding or incremental-decoder teardown.

| Field | Type | Meaning |
| --- | --- | --- |
| `error_msg` | `const char *` | Error text set when decoder status is not OK. |
| `alph_dec` | `struct ALPHDecoder *` | Optional alpha-plane decoder. |
| `alpha_data` | `const uint8_t *` | Compressed alpha payload when present. |
| `alpha_prev_line` | `const uint8_t *` | Most recently decoded alpha row or null. |

Evidence:
- Code: `thirdparty/libwebp/src/dec/vp8i_dec.h:185` — VP8Decoder

## WebPChunk

Tagged container chunk with a byte payload and linked-list successor.

Role: Represents metadata, image, frame, or container payload records.

Lifecycle: Created while parsing or editing, linked into a mux or image record, serialized during assembly, then released with mux cleanup.

| Field | Type | Meaning |
| --- | --- | --- |
| `tag` | `uint32_t` | FourCC-like chunk tag. |
| `data` | `WebPData` | Chunk payload bytes and size. |
| `next` | `WebPChunk*` | Next chunk in the list. |

Relationships:
- WebPChunk → **WebPData** (one-to-one): Each chunk contains one payload slice.
- WebPChunk → **WebPMux** (many-to-one): Many chunks are linked into one mux or one mux image record.

Evidence:
- Code: `thirdparty/libwebp/src/mux/muxi.h:39` — struct WebPChunk
- Code: `thirdparty/libwebp/src/mux/muxinternal.c` — ChunkDiskSize, MuxAssemble

## WebPConfig

Caller-supplied encoder policy and tuning state.

Role: Selects codec mode and compression trade-offs.

Lifecycle: Initialize with WebPConfigInit or WebPConfigPreset, modify fields, validate or encode with it, then retain or discard caller-owned storage.

| Field | Type | Meaning |
| --- | --- | --- |
| `lossless` | `int` | Selects lossless or lossy encoding. |
| `quality` | `float` | Controls lossy quality or lossless effort. |
| `method` | `int` | Controls quality-versus-speed trade-off. |
| `filter_strength` | `int` | Controls lossy in-loop filter strength. |
| `alpha_compression` | `int` | Selects alpha-plane compression behavior. |
| `thread_level` | `int` | Requests threaded encoding when nonzero. |

Relationships:
- WebPConfig → **WebPPicture** (one-to-one): One configuration and one picture are passed together to a WebPEncode invocation.

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPConfig, WebPConfigInit, WebPValidateConfig
- Code: `thirdparty/libwebp/src/enc/config_enc.c:27` — WebPConfigInitInternal

## WebPData

A byte slice used to exchange encoded payloads and chunk data.

Role: Carries an immutable-looking pointer-length byte range across mux and animation APIs.

Lifecycle: Caller initializes bytes and size or receives them from an API; ownership follows the specific producer or copy-data option.

| Field | Type | Meaning |
| --- | --- | --- |
| `bytes` | `const uint8_t*` | Start of byte payload. |
| `size` | `size_t` | Payload length in bytes. |

Relationships:
- WebPData → **WebPMux** (one-to-many): A mux can hold many chunk or image payload slices represented as WebPData.

Evidence:
- Code: `thirdparty/libwebp/src/webp/mux_types.h:30` — struct WebPData
- Code: `thirdparty/libwebp/src/mux/muxedit.c:189` — WebPMuxSetChunk

## WebPMux

Internal linked representation of a WebP container.

Role: Owns or references chunk lists, image records, animation parameters, and container-level state before assembly or after parsing.

Lifecycle: Create or parse a mux, add/edit images and chunks, assemble or query it, then delete it through mux API cleanup.

| Field | Type | Meaning |
| --- | --- | --- |
| `images` | `WebPMuxImage*` | Head of image or frame record list. |
| `iccp` | `WebPChunk*` | Optional ICC profile chunk list. |
| `exif` | `WebPChunk*` | Optional EXIF chunk list. |
| `xmp` | `WebPChunk*` | Optional XMP chunk list. |
| `canvas_width` | `int` | Container canvas width. |
| `canvas_height` | `int` | Container canvas height. |

Relationships:
- WebPMux → **WebPChunk** (one-to-many): One mux owns or references multiple metadata and payload chunks.
- WebPMux → **WebPMuxImage** (one-to-many): One mux chains multiple still-image or animation-frame records.

Evidence:
- Code: `thirdparty/libwebp/src/mux/muxi.h:65` — struct WebPMux
- Code: `thirdparty/libwebp/src/mux/muxinternal.c` — MuxAssemble
- Code: `thirdparty/libwebp/src/mux/muxread.c:187` — WebPMuxCreateInternal

## WebPMuxImage

Internal image or animation-frame record with header, image, alpha, and next pointers.

Role: Groups the chunks that form one still image or one animation frame.

Lifecycle: Allocated while adding or parsing image content, linked into WebPMux.images, serialized or queried, then released with its mux.

| Field | Type | Meaning |
| --- | --- | --- |
| `header` | `WebPChunk*` | Frame-header chunk when the record represents an animation frame. |
| `alpha` | `WebPChunk*` | Optional alpha chunk. |
| `img` | `WebPChunk*` | VP8 or VP8L image-data chunk. |
| `next` | `WebPMuxImage*` | Next image or frame record. |

Relationships:
- WebPMuxImage → **WebPMux** (many-to-one): Many image records can belong to one mux.
- WebPMuxImage → **WebPChunk** (one-to-many): One image record groups header, alpha, and image chunks.

Evidence:
- Code: `thirdparty/libwebp/src/mux/muxi.h:51` — struct WebPMuxImage
- Code: `thirdparty/libwebp/src/mux/muxread.c` — MuxImageCount, GetImageInfo

## WebPPicture

Public image input, output-hook, statistics, and allocation container.

Role: Carries an image through import, conversion, encoding, and optional reporting.

Lifecycle: Initialize, set dimensions and source representation, optionally allocate or import pixels, encode or transform, then free picture-owned pixel memory with WebPPictureFree.

| Field | Type | Meaning |
| --- | --- | --- |
| `use_argb` | `int` | Selects ARGB input instead of YUV(A) planes. |
| `width` | `int` | Image width in pixels. |
| `height` | `int` | Image height in pixels. |
| `y/u/v` | `uint8_t*` | Luma and chroma plane pointers. |
| `argb` | `uint32_t*` | Packed ARGB pixel-plane pointer. |
| `writer` | `WebPWriterFunction` | Optional callback used to emit compressed bytes. |
| `custom_ptr` | `void*` | Caller-defined callback context. |
| `stats` | `WebPAuxStats*` | Optional encoding statistics destination. |
| `memory_ and memory_argb_` | `void*` | Private pointers for picture-owned plane allocations. |

Relationships:
- WebPPicture → **WebPConfig** (one-to-one): A picture is encoded with one configuration per encode call.

Evidence:
- Code: `thirdparty/libwebp/src/webp/encode.h` — struct WebPPicture, WebPPictureAlloc, WebPPictureFree
- Code: `thirdparty/libwebp/src/enc/picture_enc.c:167` — WebPPictureAlloc

## WebPRescaler

Stateful fixed-point row rescaler.

Role: Accumulates source samples and produces scaled output rows across calls.

Lifecycle: Initialize for source and destination dimensions, repeatedly import and export rows, then discard caller-owned work storage after completion.

| Field | Type | Meaning |
| --- | --- | --- |
| `src_width` | `int` | Source image width. |
| `src_height` | `int` | Source image height. |
| `dst_width` | `int` | Destination image width. |
| `dst_height` | `int` | Destination image height. |
| `num_channels` | `int` | Interleaved channel count. |
| `frow` | `rescaler_t*` | Fractional working row. |
| `irow` | `rescaler_t*` | Integral working row. |
| `x_add/x_sub/y_add/y_sub` | `int` | Fixed-point horizontal and vertical stepping parameters. |

Evidence:
- Code: `thirdparty/libwebp/src/utils/rescaler_utils.h:30` — struct WebPRescaler
- Code: `thirdparty/libwebp/src/utils/rescaler_utils.c` — WebPRescalerInit, WebPRescalerImport, WebPRescalerExport

## WebRTCDataChannel

A packet-peer abstraction for data exchanged through a WebRTC peer connection.

Role: Sends and receives packet data on a named WebRTC channel.

Lifecycle: Created by a peer connection or received through a peer connection, then closed with the channel API.

Relationships:
- WebRTCDataChannel → **WebRTCPeerConnection** (many-to-one): Many channels can belong to one peer connection.

Evidence:
- Code: `modules/webrtc/webrtc_data_channel.h:35` — WebRTCDataChannel
- Code: `modules/webrtc/webrtc_peer_connection.h:35` — WebRTCPeerConnection

## WebRTCPeerConnection

A reference-counted peer-connection interface used for WebRTC negotiation and channel creation.

Role: Owns a peer-to-peer WebRTC connection abstraction.

Lifecycle: Created by application or extension code, polled during use, and closed through the peer-connection API.

Relationships:
- WebRTCPeerConnection → **WebRTCDataChannel** (one-to-many): One peer connection can create multiple data channels.

Evidence:
- Code: `modules/webrtc/webrtc_peer_connection.h:35` — WebRTCPeerConnection
- Code: `modules/webrtc/webrtc_peer_connection_extension.h` — WebRTCPeerConnectionExtension::_create_data_channel

## WebSocketMultiplayerPacket

A packet record retained by WebSocketMultiplayerPeer.

Role: Represents transport-level packet state for WebSocket multiplayer processing.

Lifecycle: Created when packet state is queued or received and discarded after multiplayer processing.

Evidence:
- Code: `modules/websocket/websocket_multiplayer_peer.h` — WebSocketMultiplayerPeer::Packet
- Code: `modules/websocket/websocket_multiplayer_peer.h:39` — WebSocketMultiplayerPeer

## XRPose

Reference-counted spatial pose published by XR tracking.

Role: Carries named pose state for an XR tracker.

Lifecycle: Created or updated by tracker services, queried by XR consumers during rendering or input work, then released through reference counting.

| Field | Type | Meaning |
| --- | --- | --- |
| `name` | `StringName` | Identifies the pose. |
| `transform` | `Transform3D` | Stores the pose transform. |
| `linear_velocity` | `Vector3` | Stores linear motion data. |
| `angular_velocity` | `Vector3` | Stores angular motion data. |
| `tracking_confidence` | `XRServer::TrackingConfidence` | Stores tracking confidence. |

Relationships:
- XRPose → **XRTracker** (many-to-one): A positional tracker can expose multiple named poses, while each pose belongs to one tracker context.

Evidence:
- Code: `servers/xr/xr_pose.h:36` — XRPose
- Code: `servers/xr/xr_positional_tracker.h:44` — XRPositionalTracker

## XRTracker

Reference-counted XR tracking source record.

Role: Classifies a tracked XR source for XRServer and positional tracker services.

Lifecycle: Created with a tracker type, registered or used by XR server services, updated by a concrete tracker implementation, and released through reference counting.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `XRServer::TrackerType` | Classifies the tracker source. |

Evidence:
- Code: `servers/xr/xr_tracker.h:40` — XRTracker
- Code: `servers/xr/xr_tracker.cpp` — XRTracker type assignment

## XrGeneratedDispatchTableCore

Generated table of OpenXR core API dispatch entries.

Role: Routes loader API entry points to the selected runtime implementation.

Lifecycle: A loader instance owns the generated dispatch table after runtime initialization and releases it when the loader instance is destroyed.

Evidence:
- Code: `thirdparty/openxr/src/xr_generated_dispatch_table_core.h:41` — struct XrGeneratedDispatchTableCore
- Code: `thirdparty/openxr/src/loader/loader_instance.hpp` — LoaderInstance::DispatchTable

## XrInstanceCreateInfo

Caller-supplied OpenXR instance-creation structure.

Role: Externally exchanged configuration for creating an OpenXR instance.

Lifecycle: The caller initializes and passes it to instance creation; the loader reads its extension chain and requested layer and extension names without taking ownership.

| Field | Type | Meaning |
| --- | --- | --- |
| `next` | `const void *` | Optional head of an extensible OpenXR structure chain. |
| `enabledApiLayerNames` | `const char * const *` | Caller-provided API-layer name array. |
| `enabledExtensionNames` | `const char * const *` | Caller-provided extension name array. |

Evidence:
- Code: `thirdparty/openxr/include/openxr/openxr.h:1202` — XrInstanceCreateInfo
- Code: `thirdparty/openxr/src/loader/loader_core.cpp:49` — xrCreateInstance

## hb_blob_t

HarfBuzz binary-data container used to hold font-file bytes and table data.

Role: font binary storage

Lifecycle: A blob is created around binary data, referenced by faces and table readers, and released through HarfBuzz object lifetime management.

| Field | Type | Meaning |
| --- | --- | --- |
| `data` | `const char *` | Pointer to the blob's binary data. |

Relationships:
- hb_blob_t → **hb_face_t** (one-to-many): A face obtains its binary font data from a blob, and a blob can be shared by face-related objects.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-blob.hh:36` — hb_blob_t
- Code: `thirdparty/harfbuzz/src/hb-face.cc` — face data blob access

## hb_buffer_t

HarfBuzz mutable text-and-glyph buffer used as shaping input and output.

Role: shaping exchange buffer

Lifecycle: A buffer is created, populated with Unicode text and segment properties, shaped with a font into glyph records and positions, then reused or destroyed.

| Field | Type | Meaning |
| --- | --- | --- |
| `info` | `hb_glyph_info_t *` | Glyph-information storage accessed during shaping. |
| `pos` | `hb_glyph_position_t *` | Glyph-position storage accessed during shaping. |
| `len` | `unsigned int` | Current number of buffer entries. |

Relationships:
- hb_buffer_t → **hb_font_t** (many-to-one): Each shaping invocation applies one font to one buffer, while a font can shape many buffers.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-buffer.hh:67` — hb_buffer_t
- Code: `thirdparty/harfbuzz/src/hb-buffer.cc` — buffer->info, buffer->pos, and buffer->len
- Code: `thirdparty/harfbuzz/src/hb-ot-shape.cc:450` — hb_ot_shape_context_t

## hb_face_t

HarfBuzz face object representing a font face backed by binary font data.

Role: font-face provider

Lifecycle: A face is built from font data, used to create fonts and shape plans, and released when no longer referenced.

Relationships:
- hb_face_t → **hb_blob_t** (many-to-one): A face accesses its OpenType file data through one backing blob.
- hb_face_t → **hb_font_t** (one-to-many): Many font instances can be created from the same face.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-face.hh:40` — hb_face_t
- Code: `thirdparty/harfbuzz/src/hb-face.cc` — OpenTypeFontFile and OpenTypeFontFace access

## hb_font_t

HarfBuzz font instance that provides glyph mapping, metrics, drawing, and variation-coordinate behavior for a face.

Role: shaping font instance

Lifecycle: A font is created from a face, configured with scale or variation state, passed to shaping, and released through HarfBuzz object lifetime management.

Relationships:
- hb_font_t → **hb_face_t** (many-to-one): Each font instance uses one face for table and glyph information.
- hb_font_t → **hb_buffer_t** (one-to-many): One font can shape many independent buffers.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-font.hh:104` — hb_font_t
- Code: `thirdparty/harfbuzz/src/hb-font.cc` — hb_font_t face and variation handling

## hb_raster_image_t

HarfBuzz raster image representation.

Role: Stores pixel-image state used by raster paint and image conversion code.

Lifecycle: Created for raster operations, passed to image helpers and paint code, then released by its owning raster operation.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-raster-image.hh:37` — hb_raster_image_t
- Code: `thirdparty/harfbuzz/src/hb-raster-paint.hh:102` — hb_raster_paint_t

## hb_shape_plan_t

HarfBuzz shaping-plan object associated with a shaping-plan key.

Role: Retains shaping configuration for shaping execution.

Lifecycle: Created for selected segment, feature, coordinate, and shaper configuration; reused or released through the shaping-plan API.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-shape-plan.hh:65` — hb_shape_plan_t
- Code: `thirdparty/harfbuzz/src/hb-shape-plan.hh:35` — hb_shape_plan_key_t

## hb_subset_accelerator_t

Subset-operation accelerator aggregate with CFF-specific accelerator declarations.

Role: Provides faster source-data access during planning and table rewriting.

Lifecycle: Associated with plan processing; its detailed ownership is not established by the supplied declaration.

| Field | Type | Meaning |
| --- | --- | --- |
| `cff1_accel` | `cff1_subset_accelerator_t` | CFF1 accelerator storage. |
| `cff2_accel` | `cff2_subset_accelerator_t` | CFF2 accelerator storage. |

Relationships:
- hb_subset_accelerator_t → **hb_subset_plan_t** (one-to-many): An accelerator type is referenced by subset plans that need source-table acceleration.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-accelerator.hh:49` — hb_subset_accelerator_t
- Code: `thirdparty/harfbuzz/src/hb-subset-plan.hh` — hb_subset_plan_t::accelerator

## hb_subset_input_t

Caller-configured selection state for a HarfBuzz font subset operation.

Role: Captures requested subset sets and input options.

Lifecycle: Configured before planning, read by subset-plan construction, and retained only for the operation that owns it.

| Field | Type | Meaning |
| --- | --- | --- |
| `sets` | `sets_t` | Groups the subset-selection sets. |

Relationships:
- hb_subset_input_t → **hb_subset_plan_t** (one-to-one): One plan records a pointer to the input that produced it.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-input.hh:84` — hb_subset_input_t
- Code: `thirdparty/harfbuzz/src/hb-subset-plan.hh` — hb_subset_plan_t::input

## hb_subset_plan_t

Derived HarfBuzz plan used while subsetting font tables.

Role: Carries selection-derived state into table subsetters.

Lifecycle: Created from an input for a subset operation, used while tables are processed, then released by the operation owner.

| Field | Type | Meaning |
| --- | --- | --- |
| `input` | `const hb_subset_input_t *` | Points to the source subset request. |
| `accelerator` | `const hb_subset_accelerator_t *` | Points to optional accelerated source-table access. |

Relationships:
- hb_subset_plan_t → **hb_subset_input_t** (one-to-one): Each plan refers to one subset input.
- hb_subset_plan_t → **hb_subset_accelerator_t** (zero-or-one-to-one): A plan may refer to one accelerator.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-subset-plan.hh:128` — hb_subset_plan_t

## hb_unicode_funcs_t

HarfBuzz Unicode-function callback object.

Role: Supplies Unicode-property operations to HarfBuzz consumers.

Lifecycle: Created and configured as a function set, then referenced by Unicode-processing operations.

Evidence:
- Code: `thirdparty/harfbuzz/src/hb-unicode.hh:40` — hb_unicode_funcs_t
- Code: `thirdparty/harfbuzz/src/hb-unicode.hh:423` — hb_unicode_range_t

## mbedtls_ssl_session

Transferable TLS session state used for session resumption and serialization.

Role: negotiated protocol session

Lifecycle: Initialized for a connection, populated during handshake, optionally copied or serialized, and freed when no longer usable.

| Field | Type | Meaning |
| --- | --- | --- |
| `tls_version` | `mbedtls_ssl_protocol_version` | Negotiated TLS protocol version. |
| `endpoint` | `int` | Client or server endpoint role. |
| `ciphersuite` | `int` | Negotiated ciphersuite identifier. |

Evidence:
- Code: `thirdparty/mbedtls/include/mbedtls/ssl.h:853` — struct mbedtls_ssl_session
- Code: `thirdparty/mbedtls/library/ssl_tls.c` — mbedtls_ssl_session_copy; mbedtls_ssl_session_save; mbedtls_ssl_session_load; mbedtls_ssl_session_free

## meshopt_Meshlet

Public meshoptimizer record that identifies the ranges and counts of one meshlet's vertices and triangles.

Role: meshlet range descriptor

Lifecycle: Caller-provided meshlet storage is populated by meshlet-building routines and later consumed with the associated vertex and triangle streams.

| Field | Type | Meaning |
| --- | --- | --- |
| `vertex_offset` | `unsigned int` | Offset into the meshlet vertex stream. |
| `triangle_offset` | `unsigned int` | Offset into the meshlet triangle stream. |
| `vertex_count` | `unsigned int` | Number of vertices in the meshlet. |
| `triangle_count` | `unsigned int` | Number of triangles in the meshlet. |

Evidence:
- Code: `thirdparty/meshoptimizer/meshoptimizer.h:701` — struct meshopt_Meshlet
- Code: `thirdparty/meshoptimizer/clusterizer.cpp` — meshlet clustering implementation
- Code: `thirdparty/meshoptimizer/meshletcodec.cpp` — meshlet stream codec

## pcg32_random_t

State for the PCG32 pseudorandom-number generator.

Role: Stores the evolving generator state and stream increment.

Lifecycle: Caller storage is seeded by `pcg32_srandom_r`, advanced by `pcg32_random_r`, and retained or discarded by the caller.

| Field | Type | Meaning |
| --- | --- | --- |
| `state` | `uint64_t` | Current internal generator state. |
| `inc` | `uint64_t` | Odd stream increment used when advancing state. |

Evidence:
- Code: `thirdparty/misc/pcg.h:11` — pcg32_random_t
- Code: `thirdparty/misc/pcg.cpp:18` — pcg32_srandom_r

## psa_key_attributes_t

PSA API object that describes the requested properties and policy of a key.

Role: key metadata and creation policy

Lifecycle: Initialized or reset by the caller, populated before generation or import, consumed by key-slot and storage code, then reset or discarded.

| Field | Type | Meaning |
| --- | --- | --- |
| `type` | `psa_key_type_t` | Key type. |
| `bits` | `psa_key_bits_t` | Key size in bits. |
| `lifetime` | `psa_key_lifetime_t` | Volatile or persistent lifetime information. |
| `id` | `mbedtls_svc_key_id_t` | Key identifier. |
| `policy` | `psa_key_policy_t` | Permitted usage and algorithm policy. |

Evidence:
- Code: `thirdparty/mbedtls/tf-psa-crypto/include/psa/crypto_struct.h:304` — struct psa_key_attributes_s
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_client.c:17` — psa_reset_key_attributes
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto.c:1762` — psa_import_key
- Code: `thirdparty/mbedtls/tf-psa-crypto/core/psa_crypto_storage.c` — persistent key storage implementation

## rcCompactHeightfield

Compact navigation spans with cells, adjacency, and area identifiers.

Role: Cross-stage traversal representation for filtering, contour extraction, regions, layers, and mesh-detail construction.

Lifecycle: Built from an `rcHeightfield`, processed by area and contour stages, then consumed while generating polygon and detail meshes.

| Field | Type | Meaning |
| --- | --- | --- |
| `cells` | `rcCompactCell *` | Grid cells indexing ranges of compact spans. |
| `spans` | `rcCompactSpan *` | Compact spans containing vertical and connectivity information. |
| `areas` | `unsigned char *` | Area identifier for each compact span. |
| `width` | `int` | Compact grid width. |
| `height` | `int` | Compact grid height. |

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Include/Recast.h:351` — rcCompactHeightfield
- Code: `thirdparty/recastnavigation/Recast/Source/RecastContour.cpp:823` — rcBuildContours

## rcHeightfield

Grid of rasterized navigation spans.

Role: Intermediate spatial representation built from source triangles for navigation-mesh construction.

Lifecycle: Recast creates the heightfield, rasterizes triangles into its spans, filters it, and converts it into a compact heightfield for later stages.

| Field | Type | Meaning |
| --- | --- | --- |
| `width` | `int` | Number of heightfield columns in the x direction. |
| `height` | `int` | Number of heightfield columns in the z direction. |
| `spans` | `rcSpan **` | Per-column linked span storage. |
| `cs` | `float` | Horizontal cell size. |
| `ch` | `float` | Vertical cell height. |

Relationships:
- rcHeightfield → **rcCompactHeightfield** (one-to-one): One source heightfield is consumed to build one compact heightfield for a build pass.

Evidence:
- Code: `thirdparty/recastnavigation/Recast/Include/Recast.h:293` — rcHeightfield
- Code: `thirdparty/recastnavigation/Recast/Source/RecastRasterization.cpp:478` — rcRasterizeTriangles

## spv::Function

SPIR-V intermediate-representation function containing control-flow blocks.

Role: shader IR function

Lifecycle: The builder creates functions while lowering shader constructs and associates blocks with each function.

| Field | Type | Meaning |
| --- | --- | --- |
| `blocks` | `std::vector<spv::Block*>` | Blocks belonging to the function. |
| `exportName` | `std::string` | Exported function name. |

Evidence:
- Code: `thirdparty/glslang/SPIRV/spvIR.h` — spv::Function

## ufbx_scene

ufbx imported-scene entity.

Role: Top-level scene type for the FBX importer API.

Lifecycle: The supplied public header indexes ufbx_scene; construction and retirement bodies are not present in the supplied excerpt.

Evidence:
- Code: `thirdparty/ufbx/ufbx.h:574` — ufbx_scene

## wslay_event_context

Wslay state object holding incoming-message slots, the active incoming/outgoing message pointers, send queues, callbacks, and frame user data.

Role: WebSocket event-processing context

Lifecycle: It persists across frame/event I/O operations and holds queued messages plus callback state for that context.

| Field | Type | Meaning |
| --- | --- | --- |
| `imsgs` | `struct wslay_event_imsg[2]` | Two incoming-message state slots. |
| `imsg` | `struct wslay_event_imsg *` | Pointer to the active incoming-message state. |
| `omsg` | `struct wslay_event_omsg *` | Pointer to the active outgoing message. |
| `send_queue` | `struct wslay_queue` | Queue of outgoing messages. |
| `send_ctrl_queue` | `struct wslay_queue` | Queue of outgoing control messages. |
| `callbacks` | `struct wslay_event_callbacks` | Application-provided event callback table. |

Evidence:
- Code: `thirdparty/wslay/wslay_event.h:80` — struct wslay_event_context
- Code: `thirdparty/wslay/wslay/wslay.h:401` — wslay_event_callbacks

## xatlas::MeshDecl

Input declaration describing mesh vertex, index, UV, normal, face-ignore, material, and polygon-size data for xatlas processing.

Role: UV-atlas mesh input

Lifecycle: The caller populates the declaration and supplies it to xatlas mesh ingestion; xatlas uses it to construct internal mesh and chart representations.

| Field | Type | Meaning |
| --- | --- | --- |
| `vertexPositionData` | `const void *` | Vertex-position data. |
| `vertexNormalData` | `const void *` | Optional vertex-normal data. |
| `vertexUvData` | `const void *` | Optional input UV data used as a chart-generation hint. |
| `indexData` | `const void *` | Optional index data. |
| `faceIgnoreData` | `const bool *` | Per-face ignore data. |
| `faceMaterialData` | `const uint32_t *` | Per-face material data. |
| `faceVertexCount` | `const uint8_t *` | Per-face polygon vertex counts. |

Evidence:
- Code: `thirdparty/xatlas/xatlas.h:109` — struct MeshDecl
- Code: `thirdparty/xatlas/xatlas.cpp` — Context and internal::Mesh
<!-- rope-ladder:end document -->
