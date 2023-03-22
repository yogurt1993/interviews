import { useCallback, useState } from "react";

import { Key } from "./Key";
import { Button } from "./Button";

import { Api } from "../utils/Api";

export const App = () => {
  const [keys, setKeys] = useState([]);
  const [isLoading, setLoadingState] = useState(false);
  const [countUsedKeys, setCountUsedKeys] = useState(0);

  const toggleLoading = useCallback(() => {
    setLoadingState(prevState=> !prevState);
  }, []);

  const incrementUsedKeys = () => {
    setCountUsedKeys((prevValue) => ++prevValue);
  }

  const decrementUsedKeys = useCallback(() => {
    setCountUsedKeys((prevValue) => --prevValue);
  }, []);

  const addKey = useCallback(async () => {
    toggleLoading();

    const key = await Api.generateKey();
    setKeys((prevKeys) => [...prevKeys, key]);

    toggleLoading();
  }, [toggleLoading]);

  const removeKey = useCallback((value) => {
    setKeys((prevKeys) => prevKeys.filter((key) => key !== value));
  }, []);

  return (
    <main>
      <div>
        <h3>Всего ключей: {keys.length}</h3>
        <h3>Использовано текущих ключей: {countUsedKeys}</h3>
      </div>

      {!keys.length && <div>Список ключей пуст</div>}

      {keys.length > 0 && (
        <div className="keys">
          {keys.map((key) => (
            <Key
              key={key}
              value={key}
              removeKey={removeKey}
              decrementUsedKeys={decrementUsedKeys}
              incrementUsedKeys={incrementUsedKeys}
            />
          ))}
        </div>
      )}

      <div>
        <Button
          onClick={addKey}
          isLoading={isLoading}
          label="Сгенерировать ключ"
        />
      </div>
    </main>
  );
};
