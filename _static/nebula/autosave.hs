behavior Autosave

    on input
        set my isChanged to true

    on blur
        send maybeSave

    on keydown debounced at 10s
        send maybeSave

    on maybeSave
        if my isChanged then 
            set my isChanged to false
            send autosave
        end
