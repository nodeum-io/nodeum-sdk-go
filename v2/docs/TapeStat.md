# TapeStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogTime** | Pointer to **string** |  | [optional] [readonly] 
**Barcode** | Pointer to **string** |  | [optional] [readonly] 
**Mounts** | Pointer to **int32** |  | [optional] [readonly] 
**DatasetsWritten** | Pointer to **int32** |  | [optional] [readonly] 
**DatasetsRead** | Pointer to **int32** |  | [optional] [readonly] 
**RecoveredWriteDataErrors** | Pointer to **int32** |  | [optional] [readonly] 
**UnrecoveredWriteDataErrors** | Pointer to **int32** |  | [optional] [readonly] 
**WriteServoErrors** | Pointer to **int32** |  | [optional] [readonly] 
**UnrecoveredWriteServoErrors** | Pointer to **int32** |  | [optional] [readonly] 
**RecoveredReadErrors** | Pointer to **int32** |  | [optional] [readonly] 
**UnrecoveredReadErrors** | Pointer to **int32** |  | [optional] [readonly] 
**LastMountUnrecoveredWriteErrors** | Pointer to **int32** |  | [optional] [readonly] 
**LastMountUnrecoveredReadErrors** | Pointer to **int32** |  | [optional] [readonly] 
**LastMountMbytesWritten** | Pointer to **int32** |  | [optional] [readonly] 
**LastMountMbytesRead** | Pointer to **int32** |  | [optional] [readonly] 
**LifetimeMbytesWritten** | Pointer to **int32** |  | [optional] [readonly] 
**LifetimeMbytesRead** | Pointer to **int32** |  | [optional] [readonly] 
**LastLoadWriteCompressionRatio** | Pointer to **int32** |  | [optional] [readonly] 
**LastLoadReadCompressionRatio** | Pointer to **int32** |  | [optional] [readonly] 
**MediumMountTime** | Pointer to **int32** |  | [optional] [readonly] 
**MediumReadyTime** | Pointer to **int32** |  | [optional] [readonly] 
**TotalNativeCapacity** | Pointer to **int32** |  | [optional] [readonly] 
**TotalUsedNativeCapacity** | Pointer to **int32** |  | [optional] [readonly] 
**WriteProtect** | Pointer to **int32** |  | [optional] [readonly] 
**Worm** | Pointer to **int32** |  | [optional] [readonly] 
**BeginningOfMediumPasses** | Pointer to **int32** |  | [optional] [readonly] 
**MiddleOfTapePasses** | Pointer to **int32** |  | [optional] [readonly] 
**ReadCompressionRatio** | Pointer to **int32** |  | [optional] [readonly] 
**WriteCompressionRatio** | Pointer to **int32** |  | [optional] [readonly] 
**MbytesTransferredToAppClient** | Pointer to **int32** |  | [optional] [readonly] 
**BytesTransferredToAppClient** | Pointer to **int32** |  | [optional] [readonly] 
**MbytesReadFromMedium** | Pointer to **int32** |  | [optional] [readonly] 
**BytesReadFromMedium** | Pointer to **int32** |  | [optional] [readonly] 
**MbytesTransferredFromAppClient** | Pointer to **int32** |  | [optional] [readonly] 
**BytesTransferredFromAppClient** | Pointer to **int32** |  | [optional] [readonly] 
**MbytesWrittenToMedium** | Pointer to **int32** |  | [optional] [readonly] 
**BytesWrittenToMedium** | Pointer to **int32** |  | [optional] [readonly] 
**DataCompressionEnabled** | Pointer to **int32** |  | [optional] [readonly] 
**WriteRetries** | Pointer to **int32** |  | [optional] [readonly] 
**WritePerms** | Pointer to **int32** |  | [optional] [readonly] 
**SuspendedWrites** | Pointer to **int32** |  | [optional] [readonly] 
**FatalSuspendedWrites** | Pointer to **int32** |  | [optional] [readonly] 
**ReadRetries** | Pointer to **int32** |  | [optional] [readonly] 
**ReadPerms** | Pointer to **int32** |  | [optional] [readonly] 
**SuspendedReads** | Pointer to **int32** |  | [optional] [readonly] 
**FatalSuspendedReads** | Pointer to **int32** |  | [optional] [readonly] 
**Partition0RemainingCapacity** | Pointer to **int32** |  | [optional] [readonly] 
**Partition1RemainingCapacity** | Pointer to **int32** |  | [optional] [readonly] 
**Partition0MaximumCapacity** | Pointer to **int32** |  | [optional] [readonly] 
**Partition1MaximumCapacity** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### GetLogTime

`func (o *TapeStat) GetLogTime() string`

GetLogTime returns the LogTime field if non-nil, zero value otherwise.

### GetLogTimeOk

`func (o *TapeStat) GetLogTimeOk() (string, bool)`

GetLogTimeOk returns a tuple with the LogTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogTime

`func (o *TapeStat) HasLogTime() bool`

HasLogTime returns a boolean if a field has been set.

### SetLogTime

`func (o *TapeStat) SetLogTime(v string)`

SetLogTime gets a reference to the given string and assigns it to the LogTime field.

### GetBarcode

`func (o *TapeStat) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *TapeStat) GetBarcodeOk() (string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBarcode

`func (o *TapeStat) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcode

`func (o *TapeStat) SetBarcode(v string)`

SetBarcode gets a reference to the given string and assigns it to the Barcode field.

### GetMounts

`func (o *TapeStat) GetMounts() int32`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *TapeStat) GetMountsOk() (int32, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMounts

`func (o *TapeStat) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMounts

`func (o *TapeStat) SetMounts(v int32)`

SetMounts gets a reference to the given int32 and assigns it to the Mounts field.

### GetDatasetsWritten

`func (o *TapeStat) GetDatasetsWritten() int32`

GetDatasetsWritten returns the DatasetsWritten field if non-nil, zero value otherwise.

### GetDatasetsWrittenOk

`func (o *TapeStat) GetDatasetsWrittenOk() (int32, bool)`

GetDatasetsWrittenOk returns a tuple with the DatasetsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetsWritten

`func (o *TapeStat) HasDatasetsWritten() bool`

HasDatasetsWritten returns a boolean if a field has been set.

### SetDatasetsWritten

`func (o *TapeStat) SetDatasetsWritten(v int32)`

SetDatasetsWritten gets a reference to the given int32 and assigns it to the DatasetsWritten field.

### GetDatasetsRead

`func (o *TapeStat) GetDatasetsRead() int32`

GetDatasetsRead returns the DatasetsRead field if non-nil, zero value otherwise.

### GetDatasetsReadOk

`func (o *TapeStat) GetDatasetsReadOk() (int32, bool)`

GetDatasetsReadOk returns a tuple with the DatasetsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatasetsRead

`func (o *TapeStat) HasDatasetsRead() bool`

HasDatasetsRead returns a boolean if a field has been set.

### SetDatasetsRead

`func (o *TapeStat) SetDatasetsRead(v int32)`

SetDatasetsRead gets a reference to the given int32 and assigns it to the DatasetsRead field.

### GetRecoveredWriteDataErrors

`func (o *TapeStat) GetRecoveredWriteDataErrors() int32`

GetRecoveredWriteDataErrors returns the RecoveredWriteDataErrors field if non-nil, zero value otherwise.

### GetRecoveredWriteDataErrorsOk

`func (o *TapeStat) GetRecoveredWriteDataErrorsOk() (int32, bool)`

GetRecoveredWriteDataErrorsOk returns a tuple with the RecoveredWriteDataErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecoveredWriteDataErrors

`func (o *TapeStat) HasRecoveredWriteDataErrors() bool`

HasRecoveredWriteDataErrors returns a boolean if a field has been set.

### SetRecoveredWriteDataErrors

`func (o *TapeStat) SetRecoveredWriteDataErrors(v int32)`

SetRecoveredWriteDataErrors gets a reference to the given int32 and assigns it to the RecoveredWriteDataErrors field.

### GetUnrecoveredWriteDataErrors

`func (o *TapeStat) GetUnrecoveredWriteDataErrors() int32`

GetUnrecoveredWriteDataErrors returns the UnrecoveredWriteDataErrors field if non-nil, zero value otherwise.

### GetUnrecoveredWriteDataErrorsOk

`func (o *TapeStat) GetUnrecoveredWriteDataErrorsOk() (int32, bool)`

GetUnrecoveredWriteDataErrorsOk returns a tuple with the UnrecoveredWriteDataErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnrecoveredWriteDataErrors

`func (o *TapeStat) HasUnrecoveredWriteDataErrors() bool`

HasUnrecoveredWriteDataErrors returns a boolean if a field has been set.

### SetUnrecoveredWriteDataErrors

`func (o *TapeStat) SetUnrecoveredWriteDataErrors(v int32)`

SetUnrecoveredWriteDataErrors gets a reference to the given int32 and assigns it to the UnrecoveredWriteDataErrors field.

### GetWriteServoErrors

`func (o *TapeStat) GetWriteServoErrors() int32`

GetWriteServoErrors returns the WriteServoErrors field if non-nil, zero value otherwise.

### GetWriteServoErrorsOk

`func (o *TapeStat) GetWriteServoErrorsOk() (int32, bool)`

GetWriteServoErrorsOk returns a tuple with the WriteServoErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWriteServoErrors

`func (o *TapeStat) HasWriteServoErrors() bool`

HasWriteServoErrors returns a boolean if a field has been set.

### SetWriteServoErrors

`func (o *TapeStat) SetWriteServoErrors(v int32)`

SetWriteServoErrors gets a reference to the given int32 and assigns it to the WriteServoErrors field.

### GetUnrecoveredWriteServoErrors

`func (o *TapeStat) GetUnrecoveredWriteServoErrors() int32`

GetUnrecoveredWriteServoErrors returns the UnrecoveredWriteServoErrors field if non-nil, zero value otherwise.

### GetUnrecoveredWriteServoErrorsOk

`func (o *TapeStat) GetUnrecoveredWriteServoErrorsOk() (int32, bool)`

GetUnrecoveredWriteServoErrorsOk returns a tuple with the UnrecoveredWriteServoErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnrecoveredWriteServoErrors

`func (o *TapeStat) HasUnrecoveredWriteServoErrors() bool`

HasUnrecoveredWriteServoErrors returns a boolean if a field has been set.

### SetUnrecoveredWriteServoErrors

`func (o *TapeStat) SetUnrecoveredWriteServoErrors(v int32)`

SetUnrecoveredWriteServoErrors gets a reference to the given int32 and assigns it to the UnrecoveredWriteServoErrors field.

### GetRecoveredReadErrors

`func (o *TapeStat) GetRecoveredReadErrors() int32`

GetRecoveredReadErrors returns the RecoveredReadErrors field if non-nil, zero value otherwise.

### GetRecoveredReadErrorsOk

`func (o *TapeStat) GetRecoveredReadErrorsOk() (int32, bool)`

GetRecoveredReadErrorsOk returns a tuple with the RecoveredReadErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecoveredReadErrors

`func (o *TapeStat) HasRecoveredReadErrors() bool`

HasRecoveredReadErrors returns a boolean if a field has been set.

### SetRecoveredReadErrors

`func (o *TapeStat) SetRecoveredReadErrors(v int32)`

SetRecoveredReadErrors gets a reference to the given int32 and assigns it to the RecoveredReadErrors field.

### GetUnrecoveredReadErrors

`func (o *TapeStat) GetUnrecoveredReadErrors() int32`

GetUnrecoveredReadErrors returns the UnrecoveredReadErrors field if non-nil, zero value otherwise.

### GetUnrecoveredReadErrorsOk

`func (o *TapeStat) GetUnrecoveredReadErrorsOk() (int32, bool)`

GetUnrecoveredReadErrorsOk returns a tuple with the UnrecoveredReadErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnrecoveredReadErrors

`func (o *TapeStat) HasUnrecoveredReadErrors() bool`

HasUnrecoveredReadErrors returns a boolean if a field has been set.

### SetUnrecoveredReadErrors

`func (o *TapeStat) SetUnrecoveredReadErrors(v int32)`

SetUnrecoveredReadErrors gets a reference to the given int32 and assigns it to the UnrecoveredReadErrors field.

### GetLastMountUnrecoveredWriteErrors

`func (o *TapeStat) GetLastMountUnrecoveredWriteErrors() int32`

GetLastMountUnrecoveredWriteErrors returns the LastMountUnrecoveredWriteErrors field if non-nil, zero value otherwise.

### GetLastMountUnrecoveredWriteErrorsOk

`func (o *TapeStat) GetLastMountUnrecoveredWriteErrorsOk() (int32, bool)`

GetLastMountUnrecoveredWriteErrorsOk returns a tuple with the LastMountUnrecoveredWriteErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastMountUnrecoveredWriteErrors

`func (o *TapeStat) HasLastMountUnrecoveredWriteErrors() bool`

HasLastMountUnrecoveredWriteErrors returns a boolean if a field has been set.

### SetLastMountUnrecoveredWriteErrors

`func (o *TapeStat) SetLastMountUnrecoveredWriteErrors(v int32)`

SetLastMountUnrecoveredWriteErrors gets a reference to the given int32 and assigns it to the LastMountUnrecoveredWriteErrors field.

### GetLastMountUnrecoveredReadErrors

`func (o *TapeStat) GetLastMountUnrecoveredReadErrors() int32`

GetLastMountUnrecoveredReadErrors returns the LastMountUnrecoveredReadErrors field if non-nil, zero value otherwise.

### GetLastMountUnrecoveredReadErrorsOk

`func (o *TapeStat) GetLastMountUnrecoveredReadErrorsOk() (int32, bool)`

GetLastMountUnrecoveredReadErrorsOk returns a tuple with the LastMountUnrecoveredReadErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastMountUnrecoveredReadErrors

`func (o *TapeStat) HasLastMountUnrecoveredReadErrors() bool`

HasLastMountUnrecoveredReadErrors returns a boolean if a field has been set.

### SetLastMountUnrecoveredReadErrors

`func (o *TapeStat) SetLastMountUnrecoveredReadErrors(v int32)`

SetLastMountUnrecoveredReadErrors gets a reference to the given int32 and assigns it to the LastMountUnrecoveredReadErrors field.

### GetLastMountMbytesWritten

`func (o *TapeStat) GetLastMountMbytesWritten() int32`

GetLastMountMbytesWritten returns the LastMountMbytesWritten field if non-nil, zero value otherwise.

### GetLastMountMbytesWrittenOk

`func (o *TapeStat) GetLastMountMbytesWrittenOk() (int32, bool)`

GetLastMountMbytesWrittenOk returns a tuple with the LastMountMbytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastMountMbytesWritten

`func (o *TapeStat) HasLastMountMbytesWritten() bool`

HasLastMountMbytesWritten returns a boolean if a field has been set.

### SetLastMountMbytesWritten

`func (o *TapeStat) SetLastMountMbytesWritten(v int32)`

SetLastMountMbytesWritten gets a reference to the given int32 and assigns it to the LastMountMbytesWritten field.

### GetLastMountMbytesRead

`func (o *TapeStat) GetLastMountMbytesRead() int32`

GetLastMountMbytesRead returns the LastMountMbytesRead field if non-nil, zero value otherwise.

### GetLastMountMbytesReadOk

`func (o *TapeStat) GetLastMountMbytesReadOk() (int32, bool)`

GetLastMountMbytesReadOk returns a tuple with the LastMountMbytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastMountMbytesRead

`func (o *TapeStat) HasLastMountMbytesRead() bool`

HasLastMountMbytesRead returns a boolean if a field has been set.

### SetLastMountMbytesRead

`func (o *TapeStat) SetLastMountMbytesRead(v int32)`

SetLastMountMbytesRead gets a reference to the given int32 and assigns it to the LastMountMbytesRead field.

### GetLifetimeMbytesWritten

`func (o *TapeStat) GetLifetimeMbytesWritten() int32`

GetLifetimeMbytesWritten returns the LifetimeMbytesWritten field if non-nil, zero value otherwise.

### GetLifetimeMbytesWrittenOk

`func (o *TapeStat) GetLifetimeMbytesWrittenOk() (int32, bool)`

GetLifetimeMbytesWrittenOk returns a tuple with the LifetimeMbytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLifetimeMbytesWritten

`func (o *TapeStat) HasLifetimeMbytesWritten() bool`

HasLifetimeMbytesWritten returns a boolean if a field has been set.

### SetLifetimeMbytesWritten

`func (o *TapeStat) SetLifetimeMbytesWritten(v int32)`

SetLifetimeMbytesWritten gets a reference to the given int32 and assigns it to the LifetimeMbytesWritten field.

### GetLifetimeMbytesRead

`func (o *TapeStat) GetLifetimeMbytesRead() int32`

GetLifetimeMbytesRead returns the LifetimeMbytesRead field if non-nil, zero value otherwise.

### GetLifetimeMbytesReadOk

`func (o *TapeStat) GetLifetimeMbytesReadOk() (int32, bool)`

GetLifetimeMbytesReadOk returns a tuple with the LifetimeMbytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLifetimeMbytesRead

`func (o *TapeStat) HasLifetimeMbytesRead() bool`

HasLifetimeMbytesRead returns a boolean if a field has been set.

### SetLifetimeMbytesRead

`func (o *TapeStat) SetLifetimeMbytesRead(v int32)`

SetLifetimeMbytesRead gets a reference to the given int32 and assigns it to the LifetimeMbytesRead field.

### GetLastLoadWriteCompressionRatio

`func (o *TapeStat) GetLastLoadWriteCompressionRatio() int32`

GetLastLoadWriteCompressionRatio returns the LastLoadWriteCompressionRatio field if non-nil, zero value otherwise.

### GetLastLoadWriteCompressionRatioOk

`func (o *TapeStat) GetLastLoadWriteCompressionRatioOk() (int32, bool)`

GetLastLoadWriteCompressionRatioOk returns a tuple with the LastLoadWriteCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastLoadWriteCompressionRatio

`func (o *TapeStat) HasLastLoadWriteCompressionRatio() bool`

HasLastLoadWriteCompressionRatio returns a boolean if a field has been set.

### SetLastLoadWriteCompressionRatio

`func (o *TapeStat) SetLastLoadWriteCompressionRatio(v int32)`

SetLastLoadWriteCompressionRatio gets a reference to the given int32 and assigns it to the LastLoadWriteCompressionRatio field.

### GetLastLoadReadCompressionRatio

`func (o *TapeStat) GetLastLoadReadCompressionRatio() int32`

GetLastLoadReadCompressionRatio returns the LastLoadReadCompressionRatio field if non-nil, zero value otherwise.

### GetLastLoadReadCompressionRatioOk

`func (o *TapeStat) GetLastLoadReadCompressionRatioOk() (int32, bool)`

GetLastLoadReadCompressionRatioOk returns a tuple with the LastLoadReadCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastLoadReadCompressionRatio

`func (o *TapeStat) HasLastLoadReadCompressionRatio() bool`

HasLastLoadReadCompressionRatio returns a boolean if a field has been set.

### SetLastLoadReadCompressionRatio

`func (o *TapeStat) SetLastLoadReadCompressionRatio(v int32)`

SetLastLoadReadCompressionRatio gets a reference to the given int32 and assigns it to the LastLoadReadCompressionRatio field.

### GetMediumMountTime

`func (o *TapeStat) GetMediumMountTime() int32`

GetMediumMountTime returns the MediumMountTime field if non-nil, zero value otherwise.

### GetMediumMountTimeOk

`func (o *TapeStat) GetMediumMountTimeOk() (int32, bool)`

GetMediumMountTimeOk returns a tuple with the MediumMountTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediumMountTime

`func (o *TapeStat) HasMediumMountTime() bool`

HasMediumMountTime returns a boolean if a field has been set.

### SetMediumMountTime

`func (o *TapeStat) SetMediumMountTime(v int32)`

SetMediumMountTime gets a reference to the given int32 and assigns it to the MediumMountTime field.

### GetMediumReadyTime

`func (o *TapeStat) GetMediumReadyTime() int32`

GetMediumReadyTime returns the MediumReadyTime field if non-nil, zero value otherwise.

### GetMediumReadyTimeOk

`func (o *TapeStat) GetMediumReadyTimeOk() (int32, bool)`

GetMediumReadyTimeOk returns a tuple with the MediumReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMediumReadyTime

`func (o *TapeStat) HasMediumReadyTime() bool`

HasMediumReadyTime returns a boolean if a field has been set.

### SetMediumReadyTime

`func (o *TapeStat) SetMediumReadyTime(v int32)`

SetMediumReadyTime gets a reference to the given int32 and assigns it to the MediumReadyTime field.

### GetTotalNativeCapacity

`func (o *TapeStat) GetTotalNativeCapacity() int32`

GetTotalNativeCapacity returns the TotalNativeCapacity field if non-nil, zero value otherwise.

### GetTotalNativeCapacityOk

`func (o *TapeStat) GetTotalNativeCapacityOk() (int32, bool)`

GetTotalNativeCapacityOk returns a tuple with the TotalNativeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalNativeCapacity

`func (o *TapeStat) HasTotalNativeCapacity() bool`

HasTotalNativeCapacity returns a boolean if a field has been set.

### SetTotalNativeCapacity

`func (o *TapeStat) SetTotalNativeCapacity(v int32)`

SetTotalNativeCapacity gets a reference to the given int32 and assigns it to the TotalNativeCapacity field.

### GetTotalUsedNativeCapacity

`func (o *TapeStat) GetTotalUsedNativeCapacity() int32`

GetTotalUsedNativeCapacity returns the TotalUsedNativeCapacity field if non-nil, zero value otherwise.

### GetTotalUsedNativeCapacityOk

`func (o *TapeStat) GetTotalUsedNativeCapacityOk() (int32, bool)`

GetTotalUsedNativeCapacityOk returns a tuple with the TotalUsedNativeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalUsedNativeCapacity

`func (o *TapeStat) HasTotalUsedNativeCapacity() bool`

HasTotalUsedNativeCapacity returns a boolean if a field has been set.

### SetTotalUsedNativeCapacity

`func (o *TapeStat) SetTotalUsedNativeCapacity(v int32)`

SetTotalUsedNativeCapacity gets a reference to the given int32 and assigns it to the TotalUsedNativeCapacity field.

### GetWriteProtect

`func (o *TapeStat) GetWriteProtect() int32`

GetWriteProtect returns the WriteProtect field if non-nil, zero value otherwise.

### GetWriteProtectOk

`func (o *TapeStat) GetWriteProtectOk() (int32, bool)`

GetWriteProtectOk returns a tuple with the WriteProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWriteProtect

`func (o *TapeStat) HasWriteProtect() bool`

HasWriteProtect returns a boolean if a field has been set.

### SetWriteProtect

`func (o *TapeStat) SetWriteProtect(v int32)`

SetWriteProtect gets a reference to the given int32 and assigns it to the WriteProtect field.

### GetWorm

`func (o *TapeStat) GetWorm() int32`

GetWorm returns the Worm field if non-nil, zero value otherwise.

### GetWormOk

`func (o *TapeStat) GetWormOk() (int32, bool)`

GetWormOk returns a tuple with the Worm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorm

`func (o *TapeStat) HasWorm() bool`

HasWorm returns a boolean if a field has been set.

### SetWorm

`func (o *TapeStat) SetWorm(v int32)`

SetWorm gets a reference to the given int32 and assigns it to the Worm field.

### GetBeginningOfMediumPasses

`func (o *TapeStat) GetBeginningOfMediumPasses() int32`

GetBeginningOfMediumPasses returns the BeginningOfMediumPasses field if non-nil, zero value otherwise.

### GetBeginningOfMediumPassesOk

`func (o *TapeStat) GetBeginningOfMediumPassesOk() (int32, bool)`

GetBeginningOfMediumPassesOk returns a tuple with the BeginningOfMediumPasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBeginningOfMediumPasses

`func (o *TapeStat) HasBeginningOfMediumPasses() bool`

HasBeginningOfMediumPasses returns a boolean if a field has been set.

### SetBeginningOfMediumPasses

`func (o *TapeStat) SetBeginningOfMediumPasses(v int32)`

SetBeginningOfMediumPasses gets a reference to the given int32 and assigns it to the BeginningOfMediumPasses field.

### GetMiddleOfTapePasses

`func (o *TapeStat) GetMiddleOfTapePasses() int32`

GetMiddleOfTapePasses returns the MiddleOfTapePasses field if non-nil, zero value otherwise.

### GetMiddleOfTapePassesOk

`func (o *TapeStat) GetMiddleOfTapePassesOk() (int32, bool)`

GetMiddleOfTapePassesOk returns a tuple with the MiddleOfTapePasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiddleOfTapePasses

`func (o *TapeStat) HasMiddleOfTapePasses() bool`

HasMiddleOfTapePasses returns a boolean if a field has been set.

### SetMiddleOfTapePasses

`func (o *TapeStat) SetMiddleOfTapePasses(v int32)`

SetMiddleOfTapePasses gets a reference to the given int32 and assigns it to the MiddleOfTapePasses field.

### GetReadCompressionRatio

`func (o *TapeStat) GetReadCompressionRatio() int32`

GetReadCompressionRatio returns the ReadCompressionRatio field if non-nil, zero value otherwise.

### GetReadCompressionRatioOk

`func (o *TapeStat) GetReadCompressionRatioOk() (int32, bool)`

GetReadCompressionRatioOk returns a tuple with the ReadCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadCompressionRatio

`func (o *TapeStat) HasReadCompressionRatio() bool`

HasReadCompressionRatio returns a boolean if a field has been set.

### SetReadCompressionRatio

`func (o *TapeStat) SetReadCompressionRatio(v int32)`

SetReadCompressionRatio gets a reference to the given int32 and assigns it to the ReadCompressionRatio field.

### GetWriteCompressionRatio

`func (o *TapeStat) GetWriteCompressionRatio() int32`

GetWriteCompressionRatio returns the WriteCompressionRatio field if non-nil, zero value otherwise.

### GetWriteCompressionRatioOk

`func (o *TapeStat) GetWriteCompressionRatioOk() (int32, bool)`

GetWriteCompressionRatioOk returns a tuple with the WriteCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWriteCompressionRatio

`func (o *TapeStat) HasWriteCompressionRatio() bool`

HasWriteCompressionRatio returns a boolean if a field has been set.

### SetWriteCompressionRatio

`func (o *TapeStat) SetWriteCompressionRatio(v int32)`

SetWriteCompressionRatio gets a reference to the given int32 and assigns it to the WriteCompressionRatio field.

### GetMbytesTransferredToAppClient

`func (o *TapeStat) GetMbytesTransferredToAppClient() int32`

GetMbytesTransferredToAppClient returns the MbytesTransferredToAppClient field if non-nil, zero value otherwise.

### GetMbytesTransferredToAppClientOk

`func (o *TapeStat) GetMbytesTransferredToAppClientOk() (int32, bool)`

GetMbytesTransferredToAppClientOk returns a tuple with the MbytesTransferredToAppClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMbytesTransferredToAppClient

`func (o *TapeStat) HasMbytesTransferredToAppClient() bool`

HasMbytesTransferredToAppClient returns a boolean if a field has been set.

### SetMbytesTransferredToAppClient

`func (o *TapeStat) SetMbytesTransferredToAppClient(v int32)`

SetMbytesTransferredToAppClient gets a reference to the given int32 and assigns it to the MbytesTransferredToAppClient field.

### GetBytesTransferredToAppClient

`func (o *TapeStat) GetBytesTransferredToAppClient() int32`

GetBytesTransferredToAppClient returns the BytesTransferredToAppClient field if non-nil, zero value otherwise.

### GetBytesTransferredToAppClientOk

`func (o *TapeStat) GetBytesTransferredToAppClientOk() (int32, bool)`

GetBytesTransferredToAppClientOk returns a tuple with the BytesTransferredToAppClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBytesTransferredToAppClient

`func (o *TapeStat) HasBytesTransferredToAppClient() bool`

HasBytesTransferredToAppClient returns a boolean if a field has been set.

### SetBytesTransferredToAppClient

`func (o *TapeStat) SetBytesTransferredToAppClient(v int32)`

SetBytesTransferredToAppClient gets a reference to the given int32 and assigns it to the BytesTransferredToAppClient field.

### GetMbytesReadFromMedium

`func (o *TapeStat) GetMbytesReadFromMedium() int32`

GetMbytesReadFromMedium returns the MbytesReadFromMedium field if non-nil, zero value otherwise.

### GetMbytesReadFromMediumOk

`func (o *TapeStat) GetMbytesReadFromMediumOk() (int32, bool)`

GetMbytesReadFromMediumOk returns a tuple with the MbytesReadFromMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMbytesReadFromMedium

`func (o *TapeStat) HasMbytesReadFromMedium() bool`

HasMbytesReadFromMedium returns a boolean if a field has been set.

### SetMbytesReadFromMedium

`func (o *TapeStat) SetMbytesReadFromMedium(v int32)`

SetMbytesReadFromMedium gets a reference to the given int32 and assigns it to the MbytesReadFromMedium field.

### GetBytesReadFromMedium

`func (o *TapeStat) GetBytesReadFromMedium() int32`

GetBytesReadFromMedium returns the BytesReadFromMedium field if non-nil, zero value otherwise.

### GetBytesReadFromMediumOk

`func (o *TapeStat) GetBytesReadFromMediumOk() (int32, bool)`

GetBytesReadFromMediumOk returns a tuple with the BytesReadFromMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBytesReadFromMedium

`func (o *TapeStat) HasBytesReadFromMedium() bool`

HasBytesReadFromMedium returns a boolean if a field has been set.

### SetBytesReadFromMedium

`func (o *TapeStat) SetBytesReadFromMedium(v int32)`

SetBytesReadFromMedium gets a reference to the given int32 and assigns it to the BytesReadFromMedium field.

### GetMbytesTransferredFromAppClient

`func (o *TapeStat) GetMbytesTransferredFromAppClient() int32`

GetMbytesTransferredFromAppClient returns the MbytesTransferredFromAppClient field if non-nil, zero value otherwise.

### GetMbytesTransferredFromAppClientOk

`func (o *TapeStat) GetMbytesTransferredFromAppClientOk() (int32, bool)`

GetMbytesTransferredFromAppClientOk returns a tuple with the MbytesTransferredFromAppClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMbytesTransferredFromAppClient

`func (o *TapeStat) HasMbytesTransferredFromAppClient() bool`

HasMbytesTransferredFromAppClient returns a boolean if a field has been set.

### SetMbytesTransferredFromAppClient

`func (o *TapeStat) SetMbytesTransferredFromAppClient(v int32)`

SetMbytesTransferredFromAppClient gets a reference to the given int32 and assigns it to the MbytesTransferredFromAppClient field.

### GetBytesTransferredFromAppClient

`func (o *TapeStat) GetBytesTransferredFromAppClient() int32`

GetBytesTransferredFromAppClient returns the BytesTransferredFromAppClient field if non-nil, zero value otherwise.

### GetBytesTransferredFromAppClientOk

`func (o *TapeStat) GetBytesTransferredFromAppClientOk() (int32, bool)`

GetBytesTransferredFromAppClientOk returns a tuple with the BytesTransferredFromAppClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBytesTransferredFromAppClient

`func (o *TapeStat) HasBytesTransferredFromAppClient() bool`

HasBytesTransferredFromAppClient returns a boolean if a field has been set.

### SetBytesTransferredFromAppClient

`func (o *TapeStat) SetBytesTransferredFromAppClient(v int32)`

SetBytesTransferredFromAppClient gets a reference to the given int32 and assigns it to the BytesTransferredFromAppClient field.

### GetMbytesWrittenToMedium

`func (o *TapeStat) GetMbytesWrittenToMedium() int32`

GetMbytesWrittenToMedium returns the MbytesWrittenToMedium field if non-nil, zero value otherwise.

### GetMbytesWrittenToMediumOk

`func (o *TapeStat) GetMbytesWrittenToMediumOk() (int32, bool)`

GetMbytesWrittenToMediumOk returns a tuple with the MbytesWrittenToMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMbytesWrittenToMedium

`func (o *TapeStat) HasMbytesWrittenToMedium() bool`

HasMbytesWrittenToMedium returns a boolean if a field has been set.

### SetMbytesWrittenToMedium

`func (o *TapeStat) SetMbytesWrittenToMedium(v int32)`

SetMbytesWrittenToMedium gets a reference to the given int32 and assigns it to the MbytesWrittenToMedium field.

### GetBytesWrittenToMedium

`func (o *TapeStat) GetBytesWrittenToMedium() int32`

GetBytesWrittenToMedium returns the BytesWrittenToMedium field if non-nil, zero value otherwise.

### GetBytesWrittenToMediumOk

`func (o *TapeStat) GetBytesWrittenToMediumOk() (int32, bool)`

GetBytesWrittenToMediumOk returns a tuple with the BytesWrittenToMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBytesWrittenToMedium

`func (o *TapeStat) HasBytesWrittenToMedium() bool`

HasBytesWrittenToMedium returns a boolean if a field has been set.

### SetBytesWrittenToMedium

`func (o *TapeStat) SetBytesWrittenToMedium(v int32)`

SetBytesWrittenToMedium gets a reference to the given int32 and assigns it to the BytesWrittenToMedium field.

### GetDataCompressionEnabled

`func (o *TapeStat) GetDataCompressionEnabled() int32`

GetDataCompressionEnabled returns the DataCompressionEnabled field if non-nil, zero value otherwise.

### GetDataCompressionEnabledOk

`func (o *TapeStat) GetDataCompressionEnabledOk() (int32, bool)`

GetDataCompressionEnabledOk returns a tuple with the DataCompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataCompressionEnabled

`func (o *TapeStat) HasDataCompressionEnabled() bool`

HasDataCompressionEnabled returns a boolean if a field has been set.

### SetDataCompressionEnabled

`func (o *TapeStat) SetDataCompressionEnabled(v int32)`

SetDataCompressionEnabled gets a reference to the given int32 and assigns it to the DataCompressionEnabled field.

### GetWriteRetries

`func (o *TapeStat) GetWriteRetries() int32`

GetWriteRetries returns the WriteRetries field if non-nil, zero value otherwise.

### GetWriteRetriesOk

`func (o *TapeStat) GetWriteRetriesOk() (int32, bool)`

GetWriteRetriesOk returns a tuple with the WriteRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWriteRetries

`func (o *TapeStat) HasWriteRetries() bool`

HasWriteRetries returns a boolean if a field has been set.

### SetWriteRetries

`func (o *TapeStat) SetWriteRetries(v int32)`

SetWriteRetries gets a reference to the given int32 and assigns it to the WriteRetries field.

### GetWritePerms

`func (o *TapeStat) GetWritePerms() int32`

GetWritePerms returns the WritePerms field if non-nil, zero value otherwise.

### GetWritePermsOk

`func (o *TapeStat) GetWritePermsOk() (int32, bool)`

GetWritePermsOk returns a tuple with the WritePerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWritePerms

`func (o *TapeStat) HasWritePerms() bool`

HasWritePerms returns a boolean if a field has been set.

### SetWritePerms

`func (o *TapeStat) SetWritePerms(v int32)`

SetWritePerms gets a reference to the given int32 and assigns it to the WritePerms field.

### GetSuspendedWrites

`func (o *TapeStat) GetSuspendedWrites() int32`

GetSuspendedWrites returns the SuspendedWrites field if non-nil, zero value otherwise.

### GetSuspendedWritesOk

`func (o *TapeStat) GetSuspendedWritesOk() (int32, bool)`

GetSuspendedWritesOk returns a tuple with the SuspendedWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuspendedWrites

`func (o *TapeStat) HasSuspendedWrites() bool`

HasSuspendedWrites returns a boolean if a field has been set.

### SetSuspendedWrites

`func (o *TapeStat) SetSuspendedWrites(v int32)`

SetSuspendedWrites gets a reference to the given int32 and assigns it to the SuspendedWrites field.

### GetFatalSuspendedWrites

`func (o *TapeStat) GetFatalSuspendedWrites() int32`

GetFatalSuspendedWrites returns the FatalSuspendedWrites field if non-nil, zero value otherwise.

### GetFatalSuspendedWritesOk

`func (o *TapeStat) GetFatalSuspendedWritesOk() (int32, bool)`

GetFatalSuspendedWritesOk returns a tuple with the FatalSuspendedWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFatalSuspendedWrites

`func (o *TapeStat) HasFatalSuspendedWrites() bool`

HasFatalSuspendedWrites returns a boolean if a field has been set.

### SetFatalSuspendedWrites

`func (o *TapeStat) SetFatalSuspendedWrites(v int32)`

SetFatalSuspendedWrites gets a reference to the given int32 and assigns it to the FatalSuspendedWrites field.

### GetReadRetries

`func (o *TapeStat) GetReadRetries() int32`

GetReadRetries returns the ReadRetries field if non-nil, zero value otherwise.

### GetReadRetriesOk

`func (o *TapeStat) GetReadRetriesOk() (int32, bool)`

GetReadRetriesOk returns a tuple with the ReadRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadRetries

`func (o *TapeStat) HasReadRetries() bool`

HasReadRetries returns a boolean if a field has been set.

### SetReadRetries

`func (o *TapeStat) SetReadRetries(v int32)`

SetReadRetries gets a reference to the given int32 and assigns it to the ReadRetries field.

### GetReadPerms

`func (o *TapeStat) GetReadPerms() int32`

GetReadPerms returns the ReadPerms field if non-nil, zero value otherwise.

### GetReadPermsOk

`func (o *TapeStat) GetReadPermsOk() (int32, bool)`

GetReadPermsOk returns a tuple with the ReadPerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadPerms

`func (o *TapeStat) HasReadPerms() bool`

HasReadPerms returns a boolean if a field has been set.

### SetReadPerms

`func (o *TapeStat) SetReadPerms(v int32)`

SetReadPerms gets a reference to the given int32 and assigns it to the ReadPerms field.

### GetSuspendedReads

`func (o *TapeStat) GetSuspendedReads() int32`

GetSuspendedReads returns the SuspendedReads field if non-nil, zero value otherwise.

### GetSuspendedReadsOk

`func (o *TapeStat) GetSuspendedReadsOk() (int32, bool)`

GetSuspendedReadsOk returns a tuple with the SuspendedReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuspendedReads

`func (o *TapeStat) HasSuspendedReads() bool`

HasSuspendedReads returns a boolean if a field has been set.

### SetSuspendedReads

`func (o *TapeStat) SetSuspendedReads(v int32)`

SetSuspendedReads gets a reference to the given int32 and assigns it to the SuspendedReads field.

### GetFatalSuspendedReads

`func (o *TapeStat) GetFatalSuspendedReads() int32`

GetFatalSuspendedReads returns the FatalSuspendedReads field if non-nil, zero value otherwise.

### GetFatalSuspendedReadsOk

`func (o *TapeStat) GetFatalSuspendedReadsOk() (int32, bool)`

GetFatalSuspendedReadsOk returns a tuple with the FatalSuspendedReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFatalSuspendedReads

`func (o *TapeStat) HasFatalSuspendedReads() bool`

HasFatalSuspendedReads returns a boolean if a field has been set.

### SetFatalSuspendedReads

`func (o *TapeStat) SetFatalSuspendedReads(v int32)`

SetFatalSuspendedReads gets a reference to the given int32 and assigns it to the FatalSuspendedReads field.

### GetPartition0RemainingCapacity

`func (o *TapeStat) GetPartition0RemainingCapacity() int32`

GetPartition0RemainingCapacity returns the Partition0RemainingCapacity field if non-nil, zero value otherwise.

### GetPartition0RemainingCapacityOk

`func (o *TapeStat) GetPartition0RemainingCapacityOk() (int32, bool)`

GetPartition0RemainingCapacityOk returns a tuple with the Partition0RemainingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartition0RemainingCapacity

`func (o *TapeStat) HasPartition0RemainingCapacity() bool`

HasPartition0RemainingCapacity returns a boolean if a field has been set.

### SetPartition0RemainingCapacity

`func (o *TapeStat) SetPartition0RemainingCapacity(v int32)`

SetPartition0RemainingCapacity gets a reference to the given int32 and assigns it to the Partition0RemainingCapacity field.

### GetPartition1RemainingCapacity

`func (o *TapeStat) GetPartition1RemainingCapacity() int32`

GetPartition1RemainingCapacity returns the Partition1RemainingCapacity field if non-nil, zero value otherwise.

### GetPartition1RemainingCapacityOk

`func (o *TapeStat) GetPartition1RemainingCapacityOk() (int32, bool)`

GetPartition1RemainingCapacityOk returns a tuple with the Partition1RemainingCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartition1RemainingCapacity

`func (o *TapeStat) HasPartition1RemainingCapacity() bool`

HasPartition1RemainingCapacity returns a boolean if a field has been set.

### SetPartition1RemainingCapacity

`func (o *TapeStat) SetPartition1RemainingCapacity(v int32)`

SetPartition1RemainingCapacity gets a reference to the given int32 and assigns it to the Partition1RemainingCapacity field.

### GetPartition0MaximumCapacity

`func (o *TapeStat) GetPartition0MaximumCapacity() int32`

GetPartition0MaximumCapacity returns the Partition0MaximumCapacity field if non-nil, zero value otherwise.

### GetPartition0MaximumCapacityOk

`func (o *TapeStat) GetPartition0MaximumCapacityOk() (int32, bool)`

GetPartition0MaximumCapacityOk returns a tuple with the Partition0MaximumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartition0MaximumCapacity

`func (o *TapeStat) HasPartition0MaximumCapacity() bool`

HasPartition0MaximumCapacity returns a boolean if a field has been set.

### SetPartition0MaximumCapacity

`func (o *TapeStat) SetPartition0MaximumCapacity(v int32)`

SetPartition0MaximumCapacity gets a reference to the given int32 and assigns it to the Partition0MaximumCapacity field.

### GetPartition1MaximumCapacity

`func (o *TapeStat) GetPartition1MaximumCapacity() int32`

GetPartition1MaximumCapacity returns the Partition1MaximumCapacity field if non-nil, zero value otherwise.

### GetPartition1MaximumCapacityOk

`func (o *TapeStat) GetPartition1MaximumCapacityOk() (int32, bool)`

GetPartition1MaximumCapacityOk returns a tuple with the Partition1MaximumCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartition1MaximumCapacity

`func (o *TapeStat) HasPartition1MaximumCapacity() bool`

HasPartition1MaximumCapacity returns a boolean if a field has been set.

### SetPartition1MaximumCapacity

`func (o *TapeStat) SetPartition1MaximumCapacity(v int32)`

SetPartition1MaximumCapacity gets a reference to the given int32 and assigns it to the Partition1MaximumCapacity field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


