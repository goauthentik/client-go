#### What's New
---

##### `POST` /stages/prompt/prompts/preview/


#### What's Changed
---

##### `GET` /flows/executor/{flow_slug}/


###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    Updated `ak-stage-prompt` component:
    * Changed property `fields` (array)

        Changed items (object):
            > Serializer for a single Prompt field

        * Changed property `type` (string)
            > * `text` - Text: Simple Text input
            > * `text_area` - Text area: Multiline Text Input.
            > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
            > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
            > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
            > * `email` - Email: Text field with Email type.
            > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
            > * `number` - Number
            > * `checkbox` - Checkbox
            > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
            > * `dropdown` - Fixed choice field rendered as a dropdown.
            > * `date` - Date
            > * `date-time` - Date Time
            > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
            > * `separator` - Separator: Static Separator Line
            > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
            > * `static` - Static: Static value, displayed as-is.
            > * `ak-locale` - authentik: Selection of locales authentik supports

##### `POST` /flows/executor/{flow_slug}/


###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    Updated `ak-stage-prompt` component:
    * Changed property `fields` (array)

        Changed items (object):
            > Serializer for a single Prompt field

        * Changed property `type` (string)
            > * `text` - Text: Simple Text input
            > * `text_area` - Text area: Multiline Text Input.
            > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
            > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
            > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
            > * `email` - Email: Text field with Email type.
            > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
            > * `number` - Number
            > * `checkbox` - Checkbox
            > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
            > * `dropdown` - Fixed choice field rendered as a dropdown.
            > * `date` - Date
            > * `date-time` - Date Time
            > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
            > * `separator` - Separator: Static Separator Line
            > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
            > * `static` - Static: Static value, displayed as-is.
            > * `ak-locale` - authentik: Selection of locales authentik supports

##### `GET` /stages/prompt/prompts/{prompt_uuid}/


###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    * Changed property `type` (string)
        > * `text` - Text: Simple Text input
        > * `text_area` - Text area: Multiline Text Input.
        > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
        > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
        > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
        > * `email` - Email: Text field with Email type.
        > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
        > * `number` - Number
        > * `checkbox` - Checkbox
        > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
        > * `dropdown` - Fixed choice field rendered as a dropdown.
        > * `date` - Date
        > * `date-time` - Date Time
        > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
        > * `separator` - Separator: Static Separator Line
        > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
        > * `static` - Static: Static value, displayed as-is.
        > * `ak-locale` - authentik: Selection of locales authentik supports

##### `PUT` /stages/prompt/prompts/{prompt_uuid}/


###### Request:

Changed content type : `application/json`

* Changed property `type` (string)
    > * `text` - Text: Simple Text input
    > * `text_area` - Text area: Multiline Text Input.
    > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
    > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
    > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
    > * `email` - Email: Text field with Email type.
    > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
    > * `number` - Number
    > * `checkbox` - Checkbox
    > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
    > * `dropdown` - Fixed choice field rendered as a dropdown.
    > * `date` - Date
    > * `date-time` - Date Time
    > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
    > * `separator` - Separator: Static Separator Line
    > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
    > * `static` - Static: Static value, displayed as-is.
    > * `ak-locale` - authentik: Selection of locales authentik supports

###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    * Changed property `type` (string)
        > * `text` - Text: Simple Text input
        > * `text_area` - Text area: Multiline Text Input.
        > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
        > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
        > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
        > * `email` - Email: Text field with Email type.
        > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
        > * `number` - Number
        > * `checkbox` - Checkbox
        > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
        > * `dropdown` - Fixed choice field rendered as a dropdown.
        > * `date` - Date
        > * `date-time` - Date Time
        > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
        > * `separator` - Separator: Static Separator Line
        > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
        > * `static` - Static: Static value, displayed as-is.
        > * `ak-locale` - authentik: Selection of locales authentik supports

##### `PATCH` /stages/prompt/prompts/{prompt_uuid}/


###### Request:

Changed content type : `application/json`

* Changed property `type` (string)
    > * `text` - Text: Simple Text input
    > * `text_area` - Text area: Multiline Text Input.
    > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
    > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
    > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
    > * `email` - Email: Text field with Email type.
    > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
    > * `number` - Number
    > * `checkbox` - Checkbox
    > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
    > * `dropdown` - Fixed choice field rendered as a dropdown.
    > * `date` - Date
    > * `date-time` - Date Time
    > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
    > * `separator` - Separator: Static Separator Line
    > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
    > * `static` - Static: Static value, displayed as-is.
    > * `ak-locale` - authentik: Selection of locales authentik supports

###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    * Changed property `type` (string)
        > * `text` - Text: Simple Text input
        > * `text_area` - Text area: Multiline Text Input.
        > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
        > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
        > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
        > * `email` - Email: Text field with Email type.
        > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
        > * `number` - Number
        > * `checkbox` - Checkbox
        > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
        > * `dropdown` - Fixed choice field rendered as a dropdown.
        > * `date` - Date
        > * `date-time` - Date Time
        > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
        > * `separator` - Separator: Static Separator Line
        > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
        > * `static` - Static: Static value, displayed as-is.
        > * `ak-locale` - authentik: Selection of locales authentik supports

##### `POST` /stages/prompt/prompts/


###### Request:

Changed content type : `application/json`

* Changed property `type` (string)
    > * `text` - Text: Simple Text input
    > * `text_area` - Text area: Multiline Text Input.
    > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
    > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
    > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
    > * `email` - Email: Text field with Email type.
    > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
    > * `number` - Number
    > * `checkbox` - Checkbox
    > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
    > * `dropdown` - Fixed choice field rendered as a dropdown.
    > * `date` - Date
    > * `date-time` - Date Time
    > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
    > * `separator` - Separator: Static Separator Line
    > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
    > * `static` - Static: Static value, displayed as-is.
    > * `ak-locale` - authentik: Selection of locales authentik supports

###### Return Type:

Changed response : **201 Created**

* Changed content type : `application/json`

    * Changed property `type` (string)
        > * `text` - Text: Simple Text input
        > * `text_area` - Text area: Multiline Text Input.
        > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
        > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
        > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
        > * `email` - Email: Text field with Email type.
        > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
        > * `number` - Number
        > * `checkbox` - Checkbox
        > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
        > * `dropdown` - Fixed choice field rendered as a dropdown.
        > * `date` - Date
        > * `date-time` - Date Time
        > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
        > * `separator` - Separator: Static Separator Line
        > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
        > * `static` - Static: Static value, displayed as-is.
        > * `ak-locale` - authentik: Selection of locales authentik supports

##### `GET` /stages/prompt/prompts/


###### Parameters:

Changed: `type` in `query`
> * `text` - Text: Simple Text input
> * `text_area` - Text area: Multiline Text Input.
> * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
> * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
> * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
> * `email` - Email: Text field with Email type.
> * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
> * `number` - Number
> * `checkbox` - Checkbox
> * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
> * `dropdown` - Fixed choice field rendered as a dropdown.
> * `date` - Date
> * `date-time` - Date Time
> * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
> * `separator` - Separator: Static Separator Line
> * `hidden` - Hidden: Hidden field, can be used to insert data into form.
> * `static` - Static: Static value, displayed as-is.
> * `ak-locale` - authentik: Selection of locales authentik supports
> 
> * `text` - Text: Simple Text input
> * `text_area` - Text area: Multiline Text Input.
> * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
> * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
> * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
> * `email` - Email: Text field with Email type.
> * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
> * `number` - Number
> * `checkbox` - Checkbox
> * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
> * `dropdown` - Fixed choice field rendered as a dropdown.
> * `date` - Date
> * `date-time` - Date Time
> * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
> * `separator` - Separator: Static Separator Line
> * `hidden` - Hidden: Hidden field, can be used to insert data into form.
> * `static` - Static: Static value, displayed as-is.
> * `ak-locale` - authentik: Selection of locales authentik supports

###### Return Type:

Changed response : **200 OK**

* Changed content type : `application/json`

    * Changed property `results` (array)

        Changed items (object):
            > Prompt Serializer

        * Changed property `type` (string)
            > * `text` - Text: Simple Text input
            > * `text_area` - Text area: Multiline Text Input.
            > * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited.
            > * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited.
            > * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames.
            > * `email` - Email: Text field with Email type.
            > * `password` - Password: Masked input, multiple inputs of this type on the same prompt need to be identical.
            > * `number` - Number
            > * `checkbox` - Checkbox
            > * `radio-button-group` - Fixed choice field rendered as a group of radio buttons.
            > * `dropdown` - Fixed choice field rendered as a dropdown.
            > * `date` - Date
            > * `date-time` - Date Time
            > * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI
            > * `separator` - Separator: Static Separator Line
            > * `hidden` - Hidden: Hidden field, can be used to insert data into form.
            > * `static` - Static: Static value, displayed as-is.
            > * `ak-locale` - authentik: Selection of locales authentik supports

